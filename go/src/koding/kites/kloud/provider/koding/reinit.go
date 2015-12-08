package koding

import (
	"koding/kites/kloud/machinestate"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"golang.org/x/net/context"
)

func (m *Machine) Reinit(ctx context.Context) (err error) {
	if err := m.UpdateState("Machine is starting", machinestate.Starting); err != nil {
		return err
	}

	// update the state to intiial state if something goes wrong, we are going
	// to change latestate to a more safe state if we passed a certain step
	// below
	latestState := m.State()
	defer func() {
		if err != nil {
			m.UpdateState("Machine is marked as "+latestState.String(), latestState)
		}
	}()

	if err := m.Session.AWSClient.Destroy(ctx, 10, 50); err != nil {
		return err
	}

	// clean up old data, so if build fails below at least we give the chance to build it again
	err = m.Session.DB.Run("jMachines", func(c *mgo.Collection) error {
		return c.UpdateId(
			m.ObjectId,
			bson.M{"$set": bson.M{
				"ipAddress":         "",
				"queryString":       "",
				"meta.instanceId":   "",
				"meta.instanceName": "",
				"status.state":      machinestate.NotInitialized.String(),
				"status.modifiedAt": time.Now().UTC(),
				"status.reason":     "Reinit cleanup",
			}},
		)
	})
	if err != nil {
		return err
	}

	// cleanup this too so "build" can continue with a clean setup
	m.IpAddress = ""
	m.QueryString = ""
	m.Meta["instance_name"] = ""
	m.Meta["instanceId"] = ""
	m.Status.State = machinestate.NotInitialized.String()

	// this updates/creates domain
	return m.Build(ctx)
}
