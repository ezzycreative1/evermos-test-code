package seeder

import (
	"fmt"
	"time"

	"01-online-store/database"
	"01-online-store/models"
)

// StatusTicketSeeder ..
func StatusTicketSeeder() {

	db := database.PostsqlConn()

	now := time.Now()

	var status []models.StatusTicket

	// DEVELOPMENT
	// var assigned = models.StatusTicket{
	// 	ID:        1,
	// 	Name:      "Assigned",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// status = append(status, assigned)

	// var closed = models.StatusTicket{
	// 	ID:        2,
	// 	Name:      "Closed",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// status = append(status, closed)

	// var progress = models.StatusTicket{
	// 	ID:        3,
	// 	Name:      "In Progress",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// status = append(status, progress)

	// var hold = models.StatusTicket{
	// 	ID:        4,
	// 	Name:      "Onhold",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// status = append(status, hold)

	// var open = models.StatusTicket{
	// 	ID:        5,
	// 	Name:      "Open",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// status = append(status, open)

	// var resolved = models.StatusTicket{
	// 	ID:        6,
	// 	Name:      "Resolved",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// status = append(status, resolved)

	// OPERASIONAL
	var assigned = models.StatusTicket{
		ID:        1,
		Name:      "Assigned",
		CreatedAt: now,
		UpdatedAt: now,
	}
	status = append(status, assigned)

	var closed = models.StatusTicket{
		ID:        2,
		Name:      "Closed",
		CreatedAt: now,
		UpdatedAt: now,
	}
	status = append(status, closed)

	var progress = models.StatusTicket{
		ID:        3,
		Name:      "In Progress",
		CreatedAt: now,
		UpdatedAt: now,
	}
	status = append(status, progress)

	var hold = models.StatusTicket{
		ID:        4,
		Name:      "Onhold",
		CreatedAt: now,
		UpdatedAt: now,
	}
	status = append(status, hold)

	var open = models.StatusTicket{
		ID:        5,
		Name:      "Open",
		CreatedAt: now,
		UpdatedAt: now,
	}
	status = append(status, open)

	var reject = models.StatusTicket{
		ID:        6,
		Name:      "Rejected",
		CreatedAt: now,
		UpdatedAt: now,
	}
	status = append(status, reject)

	var release = models.StatusTicket{
		ID:        7,
		Name:      "Released Onhold",
		CreatedAt: now,
		UpdatedAt: now,
	}
	status = append(status, release)

	var resolved = models.StatusTicket{
		ID:        8,
		Name:      "Resolved",
		CreatedAt: now,
		UpdatedAt: now,
	}
	status = append(status, resolved)

	for _, st := range status {
		if err := db.Create(&st).Error; err != nil {
			fmt.Println(err.Error())

			//log.Fatalln(err.Error())
		}
		fmt.Printf("status %s has been created\n", st.Name)
	}
}

// Priority ..
func Priority() {
	db := database.PostsqlConn()

	now := time.Now()

	var priority []models.PriorityTicket

	// OPERSIONAL
	var critical = models.PriorityTicket{
		ID:        1,
		ImpactID:  1,
		UrgencyID: 1,
		Name:      "(1) Emergency",
		CreatedAt: now,
		UpdatedAt: now,
	}
	priority = append(priority, critical)

	var major1 = models.PriorityTicket{
		ID:        2,
		ImpactID:  1,
		UrgencyID: 2,
		Name:      "(2) Major",
		CreatedAt: now,
		UpdatedAt: now,
	}
	priority = append(priority, major1)

	var hg = models.PriorityTicket{
		ID:        3,
		ImpactID:  1,
		UrgencyID: 3,
		Name:      "(3) High",
		CreatedAt: now,
		UpdatedAt: now,
	}
	priority = append(priority, hg)

	var hg2 = models.PriorityTicket{
		ID:        4,
		ImpactID:  2,
		UrgencyID: 1,
		Name:      "(2) Major",
		CreatedAt: now,
		UpdatedAt: now,
	}
	priority = append(priority, hg2)

	var low = models.PriorityTicket{
		ID:        5,
		ImpactID:  2,
		UrgencyID: 2,
		Name:      "(3) High",
		CreatedAt: now,
		UpdatedAt: now,
	}
	priority = append(priority, low)

	var medium2 = models.PriorityTicket{
		ID:        6,
		ImpactID:  2,
		UrgencyID: 3,
		Name:      "(4) Medium",
		CreatedAt: now,
		UpdatedAt: now,
	}
	priority = append(priority, medium2)

	var mj = models.PriorityTicket{
		ID:        7,
		ImpactID:  3,
		UrgencyID: 1,
		Name:      "(3) High",
		CreatedAt: now,
		UpdatedAt: now,
	}
	priority = append(priority, mj)

	var mdm = models.PriorityTicket{
		ID:        8,
		ImpactID:  3,
		UrgencyID: 2,
		Name:      "(4) Medium",
		CreatedAt: now,
		UpdatedAt: now,
	}
	priority = append(priority, mdm)

	var mdm2 = models.PriorityTicket{
		ID:        9,
		ImpactID:  3,
		UrgencyID: 3,
		Name:      "(5) Low",
		CreatedAt: now,
		UpdatedAt: now,
	}
	priority = append(priority, mdm2)

	// DEVELOPMENT
	// var critical = models.PriorityTicket{
	// 	ID:        1,
	// 	ImpactID:  1,
	// 	UrgencyID: 1,
	// 	Name:      "Critical",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// priority = append(priority, critical)

	// var major1 = models.PriorityTicket{
	// 	ID:        2,
	// 	ImpactID:  1,
	// 	UrgencyID: 2,
	// 	Name:      "Major",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// priority = append(priority, major1)

	// var hg = models.PriorityTicket{
	// 	ID:        3,
	// 	ImpactID:  1,
	// 	UrgencyID: 3,
	// 	Name:      "High",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// priority = append(priority, hg)

	// var hg2 = models.PriorityTicket{
	// 	ID:        4,
	// 	ImpactID:  2,
	// 	UrgencyID: 1,
	// 	Name:      "High",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// priority = append(priority, hg2)

	// var low = models.PriorityTicket{
	// 	ID:        5,
	// 	ImpactID:  2,
	// 	UrgencyID: 2,
	// 	Name:      "Low",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// priority = append(priority, low)

	// var medium2 = models.PriorityTicket{
	// 	ID:        6,
	// 	ImpactID:  2,
	// 	UrgencyID: 3,
	// 	Name:      "Medium",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// priority = append(priority, medium2)

	// var mj = models.PriorityTicket{
	// 	ID:        7,
	// 	ImpactID:  3,
	// 	UrgencyID: 1,
	// 	Name:      "Major",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// priority = append(priority, mj)

	// var mdm = models.PriorityTicket{
	// 	ID:        8,
	// 	ImpactID:  3,
	// 	UrgencyID: 2,
	// 	Name:      "Medium",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// priority = append(priority, mdm)

	// var mdm2 = models.PriorityTicket{
	// 	ID:        9,
	// 	ImpactID:  3,
	// 	UrgencyID: 3,
	// 	Name:      "Major",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// priority = append(priority, mdm2)

	for _, pr := range priority {
		if err := db.Create(&pr).Error; err != nil {
			fmt.Println(err.Error())

			//log.Fatalln(err.Error())
		}
		fmt.Printf("priority %s has been created\n", pr.Name)
	}
}

// TicketType ..
func TicketType() {
	db := database.PostsqlConn()

	now := time.Now()

	var ticket_type []models.TicketType

	var billing = models.TicketType{
		ID:        1,
		Name:      "billing",
		CreatedAt: now,
		UpdatedAt: now,
	}
	ticket_type = append(ticket_type, billing)

	var technical = models.TicketType{
		ID:        2,
		Name:      "technical",
		CreatedAt: now,
		UpdatedAt: now,
	}
	ticket_type = append(ticket_type, technical)

	for _, ticketType := range ticket_type {
		if err := db.Create(&ticketType).Error; err != nil {
			fmt.Println(err.Error())

			//log.Fatalln(err.Error())
		}
		fmt.Printf("ticket type %s has been created\n", ticketType.Name)
	}
}

// Impact ..
func Impact() {
	db := database.PostsqlConn()

	now := time.Now()

	var impact []models.Impact

	// OPERASIONAL
	var hgh = models.Impact{
		ID:        1,
		Name:      "(1) High",
		CreatedAt: now,
		UpdatedAt: now,
	}
	impact = append(impact, hgh)

	var md = models.Impact{
		ID:        2,
		Name:      "(2) Medium",
		CreatedAt: now,
		UpdatedAt: now,
	}
	impact = append(impact, md)

	var lw = models.Impact{
		ID:        3,
		Name:      "(3) Low",
		CreatedAt: now,
		UpdatedAt: now,
	}
	impact = append(impact, lw)

	// DEVELOPMENT
	// var hgh = models.Impact{
	// 	ID:        1,
	// 	Name:      "High",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// impact = append(impact, hgh)

	// var lw = models.Impact{
	// 	ID:        2,
	// 	Name:      "Low",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// impact = append(impact, lw)

	// var md = models.Impact{
	// 	ID:        3,
	// 	Name:      "Medium",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// impact = append(impact, md)

	for _, impact := range impact {
		if err := db.Create(&impact).Error; err != nil {
			fmt.Println(err.Error())

			//log.Fatalln(err.Error())
		}
		fmt.Printf("impact %s has been created\n", impact.Name)
	}
}

// Urgency ..
func Urgency() {
	db := database.PostsqlConn()

	now := time.Now()

	var urgency []models.Urgency

	// OPERASIONAL
	var hgh = models.Urgency{
		ID:        1,
		Name:      "(1) High",
		CreatedAt: now,
		UpdatedAt: now,
	}
	urgency = append(urgency, hgh)

	var md = models.Urgency{
		ID:        2,
		Name:      "(2) Medium",
		CreatedAt: now,
		UpdatedAt: now,
	}
	urgency = append(urgency, md)

	var lw = models.Urgency{
		ID:        3,
		Name:      "(3) Low",
		CreatedAt: now,
		UpdatedAt: now,
	}
	urgency = append(urgency, lw)

	//DEVELOPMENT
	// var hgh = models.Urgency{
	// 	ID:        1,
	// 	Name:      "High",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// urgency = append(urgency, hgh)

	// var lw = models.Urgency{
	// 	ID:        2,
	// 	Name:      "Low",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// urgency = append(urgency, lw)

	// var md = models.Urgency{
	// 	ID:        3,
	// 	Name:      "Medium",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// urgency = append(urgency, md)

	for _, urgency := range urgency {
		if err := db.Create(&urgency).Error; err != nil {
			fmt.Println(err.Error())

			//log.Fatalln(err.Error())
		}
		fmt.Printf("urgency %s has been created\n", urgency.Name)
	}
}

// RequestType ..
func RequestType() {
	db := database.PostsqlConn()

	now := time.Now()

	var rt []models.RequestType

	// OPERASIONAL

	var ir = models.RequestType{
		ID:           1,
		TicketTypeID: 1,
		Name:         "Information Request",
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	rt = append(rt, ir)

	var cr = models.RequestType{
		ID:           2,
		TicketTypeID: 2,
		Name:         "Change Request",
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	rt = append(rt, cr)

	var cc = models.RequestType{
		ID:           3,
		TicketTypeID: 2,
		Name:         "Customer Complaint",
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	rt = append(rt, cc)

	// DEVELOPMENT
	// var ad = models.RequestType{
	// 	ID:           1,
	// 	TicketTypeID: 1,
	// 	Name:         "Administration",
	// 	CreatedAt:    now,
	// 	UpdatedAt:    now,
	// }
	// rt = append(rt, ad)

	// var icd = models.RequestType{
	// 	ID:           1,
	// 	TicketTypeID: 1,
	// 	Name:         "Incident",
	// 	CreatedAt:    now,
	// 	UpdatedAt:    now,
	// }
	// rt = append(rt, icd)

	// var rfi = models.RequestType{
	// 	ID:           1,
	// 	TicketTypeID: 2,
	// 	Name:         "Request For Information",
	// 	CreatedAt:    now,
	// 	UpdatedAt:    now,
	// }
	// rt = append(rt, rfi)

	for _, rt := range rt {
		if err := db.Create(&rt).Error; err != nil {
			fmt.Println(err.Error())

			//log.Fatalln(err.Error())
		}
		fmt.Printf("request type %s has been created\n", rt.Name)
	}
}

// ServiceType ..
func ServiceType() {
	db := database.PostsqlConn()

	now := time.Now()

	var st []models.ServiceType

	var cbr = models.ServiceType{
		ID:        1,
		Name:      "Cloud Backup & Restore",
		CreatedAt: now,
		UpdatedAt: now,
	}
	st = append(st, cbr)

	var sfi = models.ServiceType{
		ID:        2,
		Name:      "Software Installation",
		CreatedAt: now,
		UpdatedAt: now,
	}
	st = append(st, sfi)

	var cpr = models.ServiceType{
		ID:        3,
		Name:      "Cloud Performance Report",
		CreatedAt: now,
		UpdatedAt: now,
	}
	st = append(st, cpr)

	var prr = models.ServiceType{
		ID:        4,
		Name:      "Password Reset Request",
		CreatedAt: now,
		UpdatedAt: now,
	}
	st = append(st, prr)

	var csca = models.ServiceType{
		ID:        5,
		Name:      "Cloud Service Configuration & Activation",
		CreatedAt: now,
		UpdatedAt: now,
	}
	st = append(st, csca)

	var ru = models.ServiceType{
		ID:        6,
		Name:      "Resource Upgrade",
		CreatedAt: now,
		UpdatedAt: now,
	}
	st = append(st, ru)

	for _, servicetype := range st {
		if err := db.Create(&servicetype).Error; err != nil {
			fmt.Println(err.Error())

			//log.Fatalln(err.Error())
		}
		fmt.Printf("service type %s has been created\n", servicetype.Name)
	}
}

// Category ..
func Category() {
	db := database.PostsqlConn()

	now := time.Now()

	var cg []models.Category

	var cs = models.Category{
		ID:        1,
		Name:      "Cloud Services",
		CreatedAt: now,
		UpdatedAt: now,
	}
	cg = append(cg, cs)

	var ns = models.Category{
		ID:        100,
		Name:      "Not Specified",
		CreatedAt: now,
		UpdatedAt: now,
	}
	cg = append(cg, ns)

	for _, cat := range cg {
		if err := db.Create(&cat).Error; err != nil {
			fmt.Println(err.Error())

			//log.Fatalln(err.Error())
		}
		fmt.Printf("category %s has been created\n", cat.Name)
	}
}

// Subcategory ..
func Subcategory() {
	db := database.PostsqlConn()

	now := time.Now()

	var scg []models.Subcategory

	var sc1 = models.Subcategory{
		ID:         1,
		CategoryID: 1,
		Name:       "Compute Services",
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	scg = append(scg, sc1)

	var sc2 = models.Subcategory{
		ID:         100,
		CategoryID: 100,
		Name:       "Not Specified",
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	scg = append(scg, sc2)

	for _, scat := range scg {
		if err := db.Create(&scat).Error; err != nil {
			fmt.Println(err.Error())

			//log.Fatalln(err.Error())
		}
		fmt.Printf("subcategory %s has been created\n", scat.Name)
	}
}

// Item ..
func Item() {
	db := database.PostsqlConn()

	now := time.Now()

	var itm []models.Item

	var itm1 = models.Item{
		ID:            1,
		CategoryID:    1,
		SubcategoryID: 1,
		Name:          "Deka Flexi",
		CreatedAt:     now,
		UpdatedAt:     now,
	}
	itm = append(itm, itm1)

	var itm2 = models.Item{
		ID:            2,
		CategoryID:    1,
		SubcategoryID: 1,
		Name:          "Deka Premium",
		CreatedAt:     now,
		UpdatedAt:     now,
	}
	itm = append(itm, itm2)

	var itm3 = models.Item{
		ID:            3,
		CategoryID:    1,
		SubcategoryID: 1,
		Name:          "Deka Prime",
		CreatedAt:     now,
		UpdatedAt:     now,
	}
	itm = append(itm, itm3)

	var itm4 = models.Item{
		ID:            4,
		CategoryID:    1,
		SubcategoryID: 1,
		Name:          "GPU Powered Compute Node",
		CreatedAt:     now,
		UpdatedAt:     now,
	}
	itm = append(itm, itm4)

	var itm100 = models.Item{
		ID:            100,
		CategoryID:    100,
		SubcategoryID: 100,
		Name:          "Not Specified",
		CreatedAt:     now,
		UpdatedAt:     now,
	}
	itm = append(itm, itm100)

	for _, item := range itm {
		if err := db.Create(&item).Error; err != nil {
			fmt.Println(err.Error())

			//log.Fatalln(err.Error())
		}
		fmt.Printf("item %s has been created\n", item.Name)
	}
}

// Site ..
func Site() {
	db := database.PostsqlConn()

	now := time.Now()

	var site []models.Site

	var st1 = models.Site{
		ID:        1,
		Name:      "TBS - Lintasarta Building",
		CreatedAt: now,
		UpdatedAt: now,
	}
	site = append(site, st1)

	var st2 = models.Site{
		ID:        2,
		Name:      "Technopark - Serpong BSD",
		CreatedAt: now,
		UpdatedAt: now,
	}
	site = append(site, st2)

	var st3 = models.Site{
		ID:        3,
		Name:      "SB - Jatiluhur Purwakarta",
		CreatedAt: now,
		UpdatedAt: now,
	}
	site = append(site, st3)

	//DEVELOPMENT
	// var st4 = models.Site{
	// 	ID:        1,
	// 	Name:      "Jakarta",
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }
	// site = append(site, st4)

	for _, st := range site {
		if err := db.Create(&st).Error; err != nil {
			fmt.Println(err.Error())

			//log.Fatalln(err.Error())
		}
		fmt.Printf("site %s has been created\n", st.Name)
	}
}
