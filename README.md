# Go Backend for ProFitnessDeals.com

- [Description](#description)
- [Features](#features)
- [License](#license)
- [Installation](#installation)
- [Contact Information](#contact-information)

## Description

Go backend managing the server-side functionality of profitnessdeals. You can find the Nextjs(TypeScript) [frontend](https://github.com/michaelhunter555/profitnessdeals) here to see how it all comes together. This app retrieves data from Amazon's Produt Advertising Api V5 and product information for the profitnessdeals client. Prices are dynamic and therefore advertisable without violation of Amazon's terms.

## Features

- jwt token authentication & refresh
- protected auth routes
- amazon product advertising api v5
- mongodb as noSQL database

## License

- Distributed under the MIT License. See `LICENSE` for more information.

## Installation

1. Clone the repository:
   ```sh
   go get github.com/michaelhunter555/gofitnessdeals-backend
   ```

### Environment Variables

- Create a `.env` file in the root of your project and add the following:
  ```
  MONGO_DB_URI=mongodb+srv://<username>:<password>@<clusterUrl>/
  JWT_SECRET= /*create in bash entering: openssl rand -base64 32 /*
  ```

AMAZON_ACCESS_KEY=AmazonAcessKey
AMAZON_SECRET_KEY=AmazonSecretKey

```

## Contact Information

- Michael Hunter - michaelhunterbkk@gmail.com
```
