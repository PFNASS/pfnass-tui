package app

type serviceGenerator struct {
	titles     []string
	descs      []string
	titleIndex int
	descIndex  int
}

func (r *serviceGenerator) set() {
	r.titles = []string{
		"Ansible Playbooks",
		"Docker Network",
		"K3s",
		"Home Assistant",
		"PFRPS",
	}
	r.descs = []string{
		"Ansible Playbooks for automation",
		"Docker network setup for containers",
		"K3s lightweight Kubernetes cluster",
		"Home Assistant for home automation",
		"PFRPS Server",
	}
	r.titleIndex = 0
	r.descIndex = 0
}

func (r *serviceGenerator) next() item {
	if r.titleIndex == 0 || r.descIndex == 0 {
		r.set()
	}

	i := item{
		title:       r.titles[r.titleIndex],
		description: r.descs[r.descIndex],
	}

	r.titleIndex++
	if r.titleIndex >= len(r.titles) {
		r.titleIndex = 0
	}

	r.descIndex++
	if r.descIndex >= len(r.descs) {
		r.descIndex = 0
	}

	return i
}