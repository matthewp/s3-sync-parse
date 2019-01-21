package s3syncparse

import (
	"testing"
)

func example() string {
	examp := `upload: public/baked-flounder-with-tomatoes-and-basil/index.html to s3://fish.recipes/baked-flounder-with-tomatoes-and-basil/index.html
	upload: public/sweet-chili-glazed-salmon/index.html to s3://fish.recipes/sweet-chili-glazed-salmon/index.html
	upload: public/salmon-with-red-pepper-pesto/index.html to s3://fish.recipes/salmon-with-red-pepper-pesto/index.html
	upload: public/smoked-salmon-pizza/index.html to s3://fish.recipes/smoked-salmon-pizza/index.html
	upload: public/honey-mustard-catfish/index.html to s3://fish.recipes/honey-mustard-catfish/index.html
	upload: public/chipotle-lime-salmon/index.html to s3://fish.recipes/chipotle-lime-salmon/index.html
	upload: public/herb-crusted-salmon/index.html to s3://fish.recipes/herb-crusted-salmon/index.html
	upload: public/swordfish-peppercorn-butter-sauce/index.html to s3://fish.recipes/swordfish-peppercorn-butter-sauce/index.html
	upload: public/cilantro-lime-tilapia-tacos/index.html to s3://fish.recipes/cilantro-lime-tilapia-tacos/index.html
	upload: public/salmon-skewers-with-cucumber/index.html to s3://fish.recipes/salmon-skewers-with-cucumber/index.html
	upload: public/citrus-mahi-mahi-with-brown-rice-noodles/index.html to s3://fish.recipes/citrus-mahi-mahi-with-brown-rice-noodles/index.html
	upload: public/steamed-flounder-with-grape-tomato-sauce/index.html to s3://fish.recipes/steamed-flounder-with-grape-tomato-sauce/index.html`

	return examp
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func TestParse(t *testing.T) {
	u := ParseAll(example())

	assertEqual(t, u[0].Src, "public/baked-flounder-with-tomatoes-and-basil/index.html")
	assertEqual(t, u[0].Dest, "s3://fish.recipes/baked-flounder-with-tomatoes-and-basil/index.html")
	assertEqual(t, u[0].Path, "/baked-flounder-with-tomatoes-and-basil/index.html")
}
