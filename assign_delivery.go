package main

func assign_order() {

	var available_agent *deliveryagent

	for _, agent := range agentData {
		if agent.is_available {
			available_agent = agent
			break
		}
	}

}
