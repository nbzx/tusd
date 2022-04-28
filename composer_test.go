package tusd_test

import (
	"github.com/nbzx/tusd"
	"github.com/nbzx/tusd/filestore"
	"github.com/nbzx/tusd/limitedstore"
	"github.com/nbzx/tusd/memorylocker"
)

func ExampleNewStoreComposer() {
	composer := tusd.NewStoreComposer()

	fs := filestore.New("./data")
	fs.UseIn(composer)

	ml := memorylocker.New()
	ml.UseIn(composer)

	ls := limitedstore.New(1024*1024*1024, composer.Core, composer.Terminater)
	ls.UseIn(composer)

	config := tusd.Config{
		StoreComposer: composer,
	}

	_, _ = tusd.NewHandler(config)
}
