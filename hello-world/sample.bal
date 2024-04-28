import ballerina/http;
import ballerina/os;
import ballerina/io;

service / on new http:Listener(8090) {
    resource function get .() returns json|error {
        string serviceURL = os:getEnv("SVC_URL");
        string tokenURL = os:getEnv("TOKEN_URL");
        string consumerKey = os:getEnv("CONSUMER_KEY");
        string consumerSecret = os:getEnv("CONSUMER_SECRET");
        io:println("url:"+serviceURL);
        io:println("TOKEN_URL:"+tokenURL);
        io:println("CONSUMER_KEY"+consumerKey);
        io:println("CONSUMER_SECRET:"+consumerSecret);

        http:Client httpClient = check new http:Client(serviceURL, {
            timeout: 60,
            auth: {
                tokenUrl: tokenURL,
                clientId: consumerKey,
                clientSecret: consumerSecret
            }
        });
    

        json payload = check httpClient->get("/");

        
            return payload;
    }
    
}
