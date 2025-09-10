import ballerina/http;

type Greeting record {
    string 'from;
    string to;
    string message;
};

service / on new http:Listener(8090) {
    resource function get greeting() returns Greeting {
        Greeting greetingMessage = {"from" : "Choreo", "to" : "hansii", "message" : "Welcome to Choreo v1.2!"};
        return greetingMessage;
    }

    resource function get test() returns Greeting {
        Greeting greetingMessage = {"from" : "Choreo", "to" : "hansi", "message" : "Welcome to Choreo!"};
        return greetingMessage;
    }
}
