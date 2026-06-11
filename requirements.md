## Open source Zipnova client go library

### Requirements

The main idea of this project is to create a golang library for other projects to import and be able to integrate "Zipnova" which is a third party shipping management service for Argentina

For now, the library will provide functionality to calculate shipping prices, create shipping orgers and track them.


### Zipnova api:

#### Mandatory headers

**Required Headers**

All API responses that return data will use JSON as their response format.

For `POST` or `PUT` requests that send data, the request body must also use JSON.

When building the request, you must include the following header (in addition to the authentication header):

```http
Accept: application/json

```

If you are sending data to Zipnova in the request body, you must also include the following header:

```http
Content-Type: application/json

```


#### Auth
Your requests must use **HTTP Basic Authentication**.

-   **Username:** API Token
-   **Password:** API Secret

There should be an mcp installed with zipnova docs, please implement this go library providing functionalities to  calculate shipping prices, create shipping orgers and track them.
