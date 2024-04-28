# samples-for-connections-testing

These repos are used to test connections

| Component Type | Repo Name |
| --- | ----------- |
| **Manual Task**
| buildpackJob | hello-world-go-task |
| byocJob | hello-world-go-task  |
| manualTrigger | ballerina-task |
| miJob | weather-to-logs-mi-manual-task |
| **Scheduled Task**
| buildpackCronJob | hello-world-go-task |
| byocCronJob | hello-world-go-task  |
| scheduledTask | ballerina-task |
| miCronJob | weather-to-logs-mi-manual-task |
| **Webhooks**
| buildpackWebhook | greeting-service-go |
| byockWebhook | greeting-service-go  |
| webhook | hello-world |
| miWebhook | hello-world-mi|
| **Event Handlers**
| buildpackEventHandler | containerized-rabbitmq-listener |
| byocEventHandler | containerized-rabbitmq-listener  |
| ballerinaEventHandler | ballerina-task ( can test with a manual task )|
| miEventHandler | mi-rabbitmq-listener|
| **Test Runners**
| buildpackTestRunner | test-runner-go |
| byocTestRunner | test-runner-go |
| **Services**
| miApiService | hello-world-mi |
| ballerinaService | hello-world |


### Note
- ```greeting-service``` repo is used as the dependent service (publisher component) in all subscriber components to create the connection. 

- Update the .choreo/component-config.yaml file in each repo to update connection configurations accordingly

- Follow these [ steps ](https://docs.google.com/document/d/1Ua8UA2bhp5pg9sFk25F2B9JwC53Egx3BcwCPz9nxYDc/edit#heading=h.5tfsdp5hr3u2) to setup the rabbitMQ instance for testing event handlers

