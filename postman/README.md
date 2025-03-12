# URL Shortener API - Postman Collection

This directory contains a Postman collection for testing the URL Shortener API.

## Getting Started

1. Install [Postman](https://www.postman.com/downloads/) if you haven't already.
2. Import the collection file `url-shortener-api.json` into Postman.
3. Set up your environment variables.

## Environment Variables

The collection uses the following environment variables:

| Variable    | Description                                   | Default Value                                        |
|-------------|-----------------------------------------------|------------------------------------------------------|
| BASE_URL    | The base URL of the API                       | http://localhost:8080                                |
| API_KEY     | Your API key for authentication               | your_secure_api_key_here                             |
| SHORT_PATH  | An example short path for testing             | a1b2c3d4                                             |
| LONG_URL    | An example long URL to be shortened           | https://example.com/very/long/url/that/needs/shortening |

## Setting Up Environment Variables in Postman

1. In Postman, click on the "Environments" tab in the sidebar.
2. Click the "+" button to create a new environment.
3. Name it "URL Shortener API".
4. Add the variables listed above with your specific values.
5. Click "Save".
6. Select the "URL Shortener API" environment from the environment dropdown in the top right corner.

## Authentication

The collection includes two authentication methods:

1. **API Key Authentication** - Uses the `X-API-Key` header with your API key.
2. **Bearer Token Authentication** - Uses the `Authorization` header with `Bearer {your_api_key}`.

Both methods are supported by the API and are included in the collection for demonstration purposes.

## Available Endpoints

### Public Endpoints

- **Health Check** - `GET /api/health`
  - Check if the API is up and running.
  - No authentication required.

- **Redirect to Original URL** - `GET /s/{shortPath}`
  - Redirect to the original URL associated with the given short path.
  - No authentication required.

### Protected Endpoints

- **Shorten URL** - `POST /api/shorten`
  - Create a shortened URL for a given long URL.
  - Requires authentication.
  - Request body: `{ "url": "https://example.com/long/url" }`

- **Get URL Statistics** - `GET /api/stats/{shortPath}`
  - Get statistics for a shortened URL.
  - Requires authentication.

## Example Workflow

1. Use the "Health Check" request to verify the API is running.
2. Use the "Shorten URL" request to create a new short URL.
3. Copy the `short_path` from the response.
4. Update the `SHORT_PATH` environment variable with the copied value.
5. Use the "Get URL Statistics" request to view statistics for the URL.
6. Use the "Redirect to Original URL" request to test the redirection.

## Response Examples

The collection includes example responses for each endpoint, showing both successful and error scenarios. 
