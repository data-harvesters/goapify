<h1  align="center">GoApify</h1>
<div align="center">
  <strong> a golang apify package.</strong>
</div>
<br />

  

## About
[Apify](https://apify.com) is a platform to deploy and publish web scrappers and web automation tools. This package aims to bring a simple yet powerful Golang package to interact with Apify's Actor API. 

## Usage
```go
    type input struct {
        Test string `json:"test"`
    }

    a := goapify.NewActor(
		os.Getenv("APIFY_DEFAULT_KEY_VALUE_STORE_ID"),
		os.Getenv("APIFY_TOKEN"),
		os.Getenv("APIFY_DEFAULT_DATASET_ID"),
	)

	i := new(input)

	err := a.Input(i)
	if err != nil {
		panic(err)
	}

```
