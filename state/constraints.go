// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"fmt"

	"github.com/juju/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"

	"github.com/juju/juju/core/constraints"
	"github.com/juju/juju/core/instance"
)

type constraintsWithID struct {
	DocID  string         `bson:"_id"`
	Nested constraintsDoc `bson:",inline"`
}

// constraintsDoc is the mongodb representation of a constraints.Value.
type constraintsDoc struct {
	ModelUUID      string `bson:"model-uuid"`
	Arch           *string
	CpuCores       *uint64
	CpuPower       *uint64
	Mem            *uint64
	RootDisk       *uint64
	RootDiskSource *string
	InstanceType   *string
	Container      *instance.ContainerType
	Tags           *[]string
	Spaces         *[]string
	VirtType       *string
	Zones          *[]string
}

func (doc constraintsDoc) value() constraints.Value {
	result := constraints.Value{
		Arch:           doc.Arch,
		CpuCores:       doc.CpuCores,
		CpuPower:       doc.CpuPower,
		Mem:            doc.Mem,
		RootDisk:       doc.RootDisk,
		RootDiskSource: doc.RootDiskSource,
		InstanceType:   doc.InstanceType,
		Container:      doc.Container,
		Tags:           doc.Tags,
		Spaces:         doc.Spaces,
		VirtType:       doc.VirtType,
		Zones:          doc.Zones,
	}
	return result
}

func newConstraintsDoc(cons constraints.Value) constraintsDoc {
	result := constraintsDoc{
		Arch:           cons.Arch,
		CpuCores:       cons.CpuCores,
		CpuPower:       cons.CpuPower,
		Mem:            cons.Mem,
		RootDisk:       cons.RootDisk,
		RootDiskSource: cons.RootDiskSource,
		InstanceType:   cons.InstanceType,
		Container:      cons.Container,
		Tags:           cons.Tags,
		Spaces:         cons.Spaces,
		VirtType:       cons.VirtType,
		Zones:          cons.Zones,
	}
	return result
}

func createConstraintsOp(id string, cons constraints.Value) txn.Op {
	return txn.Op{
		C:      constraintsC,
		Id:     id,
		Assert: txn.DocMissing,
		Insert: newConstraintsDoc(cons),
	}
}

func setConstraintsOp(id string, cons constraints.Value) txn.Op {
	return txn.Op{
		C:      constraintsC,
		Id:     id,
		Assert: txn.DocExists,
		Update: bson.D{{"$set", newConstraintsDoc(cons)}},
	}
}

func removeConstraintsOp(id string) txn.Op {
	return txn.Op{
		C:      constraintsC,
		Id:     id,
		Remove: true,
	}
}

func readConstraints(mb modelBackend, id string) (constraints.Value, error) {
	constraintsCollection, closer := mb.db().GetCollection(constraintsC)
	defer closer()

	doc := constraintsDoc{}
	if err := constraintsCollection.FindId(id).One(&doc); err == mgo.ErrNotFound {
		return constraints.Value{}, errors.NotFoundf("constraints")
	} else if err != nil {
		return constraints.Value{}, err
	}
	return doc.value(), nil
}

func writeConstraints(mb modelBackend, id string, cons constraints.Value) error {
	ops := []txn.Op{setConstraintsOp(id, cons)}
	if err := mb.db().RunTransaction(ops); err != nil {
		return fmt.Errorf("cannot set constraints: %v", err)
	}
	return nil
}

// AllConstraints retrieves all the constraints in the model
// and provides a way to query based on machine or application.
func (m *Model) AllConstraints() (*ModelConstraints, error) {
	coll, closer := m.st.db().GetCollection(constraintsC)
	defer closer()

	var docs []constraintsWithID
	err := coll.Find(nil).All(&docs)
	if err != nil {
		return nil, errors.Annotate(err, "cannot get all constraints for model")
	}
	all := &ModelConstraints{
		data: make(map[string]constraintsDoc),
	}
	for _, doc := range docs {
		all.data[m.localID(doc.DocID)] = doc.Nested
	}
	return all, nil
}

// ModelConstraints represents all the contraints in  a model
// keyed on global key.
type ModelConstraints struct {
	data map[string]constraintsDoc
}

// Machine returns the constraints for the specified machine.
func (c *ModelConstraints) Machine(machineID string) constraints.Value {
	doc, found := c.data[machineGlobalKey(machineID)]
	if !found {
		return constraints.Value{}
	}
	return doc.value()
}

// TODO: add methods for Model and Application constraints checks
// if desired. Not currently used in status calls.
