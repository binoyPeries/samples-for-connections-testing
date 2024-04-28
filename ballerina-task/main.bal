
import ballerina/io;
import ballerina/http;
import ballerina/os;

public function main() returns error?{

       string serviceURL = os:getEnv("SVC_URL");
        string tokenURL = os:getEnv("TOKEN_URL");
        string consumerKey = os:getEnv("CONSUMER_KEY");
        string consumerSecret = os:getEnv("CONSUMER_SECRET");
        
        http:Client httpClient = check new http:Client(serviceURL, {
            auth: {
                tokenUrl: tokenURL,
                clientId: consumerKey,
                clientSecret: consumerSecret
            }
        });
    

        json payload = check httpClient->get("/");

        
    io:println(payload);
}
