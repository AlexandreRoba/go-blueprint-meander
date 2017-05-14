package meander

type j struct {
	Name       string
	PlaceTypes []string
}

var Journeys = []interface{}{
	j{
		Name: "Romantic",
		PlaceTypes: []string{
			"park",
			"bar",
			"movie_theatre",
			"restaurant",
			"florist",
			"taxi_stand",
		},
	},
	j{
		Name: "Shopping",
		PlaceTypes: []string{
			"department_store",
			"cafe",
			"clothing_store",
			"jewelry_store",
			"shoe_store",
		},
	},
	j{
		Name: "Night Out",
		PlaceTypes: []string{
			"bar",
			"casino",
			"food",
			"night_club",
			"bar",
			"bar",
			"hostpital",
		},
	},
	j{
		Name: "Culture",
		PlaceTypes: []string{
			"museum",
			"cafe",
			"cemetery",
			"library",
			"art_gallery",
		},
	},
	j{
		Name: "Pamper",
		PlaceTypes: []string{
			"hair_care",
			"beaty_Salon",
			"cafe",
			"spa",
		},
	},

}
