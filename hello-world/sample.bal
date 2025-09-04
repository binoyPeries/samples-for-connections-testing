import ballerina/http;
import ballerina/os;
import ballerina/io;

service /hello on new http:Listener(8090) {

    resource function get greeting() returns json|error {
        // Load configs from environment
        string serviceURL = os:getEnv("SVC_URL");
        string tokenURL = os:getEnv("TOKEN_URL");
        string consumerKey = os:getEnv("CONSUMER_KEY");
        string consumerSecret = os:getEnv("CONSUMER_SECRET");

        io:println("SVC_URL: " + serviceURL);
        io:println("TOKEN_URL: " + tokenURL);
        io:println("CONSUMER_KEY: " + consumerKey);
        io:println("CONSUMER_SECRET: " + consumerSecret);

        // Create HTTP client with OAuth2 client credentials
        http:Client httpClient = check new (serviceURL, {
            timeout: 60,
            auth: {
                tokenUrl: tokenURL,
                clientId: consumerKey,
                clientSecret: consumerSecret
            }
        });

        // GraphQL query payload
        json gqlPayload = {
            query: "query GreetWorld { greeting(name: \"World\") }"
        };

        // Send request to GraphQL server
        http:Response resp = check httpClient->post("", gqlPayload);

        // Extract response JSON
        json result = check resp.getJsonPayload();

        return result;
    }
}
