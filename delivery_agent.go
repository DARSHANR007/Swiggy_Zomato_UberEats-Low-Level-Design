package main

type deliveryagent struct {
	Name                string
	deliveryagent_id    int
	phone               int
	vehicle_number      string
	deliveryagent_email string
	is_available        bool
	assigned_orders     []int
}

var agentData map[int]*deliveryagent

func register_Agent(Name string, deliveryagent_id int, phone int, vehicle_number string, deliveryagent_email string, is_available bool, assigned_orders []int) *deliveryagent {

	new_user := &deliveryagent{
		Name:                Name,
		deliveryagent_id:    deliveryagent_id,
		phone:               phone,
		vehicle_number:      vehicle_number,
		deliveryagent_email: deliveryagent_email,
		is_available:        false,
		assigned_orders:     assigned_orders,
	}

	agentData[deliveryagent_id] = new_user
	return new_user

}

func (this *deliveryagent) GetAgentData(id int) *deliveryagent {
	return agentData[id]
}
