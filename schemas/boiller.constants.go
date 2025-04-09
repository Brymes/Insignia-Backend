package schemas

type PropertyType string
type BoilerFuelType string
type BoilerType string
type BoilerAge string
type BoilerMounting string
type BedroomCount string
type BathroomCount string
type UserType string

const (
	// Property Types
	DetachedHouse     PropertyType = "Detached house"
	SemiDetachedHouse PropertyType = "Semi-detached house"
	Terrace           PropertyType = "Terrace"
	Flat              PropertyType = "Flat"
	Bungalow          PropertyType = "Bungalow"

	// Boiler Fuel Types
	Gas  BoilerFuelType = "Gas"
	LPG  BoilerFuelType = "LPG"
	Oil  BoilerFuelType = "Oil"
	None BoilerFuelType = "N/A"

	// Boiler Types
	Combi      BoilerType = "Combi"
	System     BoilerType = "System"
	BackBoiler BoilerType = "Back boiler"
	NABoiler   BoilerType = "N/A"

	// Boiler Age
	Age0To10   BoilerAge = "0-10 years"
	Age10To20  BoilerAge = "10-20 years"
	Age20To30  BoilerAge = "20-30 years"
	AgeUnknown BoilerAge = "N/A"

	// Boiler Mounting
	WallMounted   BoilerMounting = "Wall-mounted"
	FloorStanding BoilerMounting = "Floor-standing"
	NAMounting    BoilerMounting = "N/A"

	// Bedroom Count
	Bedrooms1To2  BedroomCount = "1-2"
	Bedrooms2To4  BedroomCount = "2-4"
	Bedrooms4To6  BedroomCount = "4-6"
	Bedrooms6Plus BedroomCount = "6+"

	// Bathroom Count
	Bathrooms1To2  BathroomCount = "1-2"
	Bathrooms2To4  BathroomCount = "2-4"
	Bathrooms4Plus BathroomCount = "4+"

	// User Type
	Homeowner UserType = "Homeowner"
	Landlord  UserType = "Landlord"
)
