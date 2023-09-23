# Service for a dynamic segmentation of users
This service stores the users and the segments to which they belong, allowing the creation and deletion of segments, adding and removing users from a segment, and also receiving a report for a certain period with the history of additions and removals of a certain user from segments.

### Running the service:
- Clone the repository
- In the project repository
    ```bash
    docker-compose up --build
    ```
- The service is running on `localhost:8080`
- Swagger documentation is available at `http://localhost:8080/docs/index.html`

### Running tests:
- Clone the repository
- In the project repository
    ```bash
    docker-compose --file docker-compose_test.yml up --build --abort-on-container-exit
    ```

### Requests examples:
1. Method for creating a segment.
    - Request: 
        ```bash
        curl --header "Content-Type: application/json" \
             --request POST \
             --data '{"name":"VOICE_MESSAGES"}' \
             http://localhost:8080/createSegment
        ```
    - Response:
        ```bash
        HTTP/1.1 201 Created
        Content-Type: application/json
        ...
        {"name":"VOICE_MESSAGES"}
        ```
2. Method for deleting a segment. Accepts the segment name (`name`). 
    - Request: 
        ```bash
        curl --header "Content-Type: application/json" \
             --request DELETE \
             --data '{"name":"VOICE_MESSAGES"}' \
             http://localhost:8080/deleteSegment
        ```
    - Response:
        ```bash
        HTTP/1.1 200 OK
        ```
3. Method for adding a user to a segment. Accepts a list of segment names to add to the user (`segmentsToAdd`), a list of segment names to remove from the user (`segmentsToDelete`) and the user ID (`userId`).
    - Request: 
        ```bash
        curl --header "Content-Type: application/json" \
             --request POST \
             --data '{"segmentsToAdd":["VOICE_MESSAGES"], "segmentsToDelete": ["DISCOUNT_30"], "userId": 1}' \
             http://localhost:8080/updateUserSegments
        ```
    - Response:
        ```bash
        HTTP/1.1 201 CREATED
        ```
4. Method for getting a user's active segments. Accepts the user ID (`userId`).
    - Request: 
        ```bash
        curl http://localhost:8080/userSegments?userId=1
        ```
    - Response:
        ```bash
        HTTP/1.1 200 OK
        Content-Type: application/json
        ...
        ["VOICE_MESSAGES"]
        ```
5. Method for generating a report on a user's segment history. Accepts the user ID (`userId`), year (`year`) and month (`month`).
    - Request:
        ```bash
        curl "http://localhost:8080/generateReport?year=2023&month=9&userId=1"
        ```
    - Response:
        ```bash
        HTTP/1.1 200 OK
        Content-Type: application/json
        ...
        {"link":"reports/9849d4bb-cf38-491b-a45a-b83af811046c.csv"}
        ```
6. Method for getting a report. Accepts the `link`  from the response to the previous request.
    - Request: 
        ```bash
        curl "http://localhost:8080/reports/9849d4bb-cf38-491b-a45a-b83af811046c.csv"
        ```
    - Response:
        ```bash
        HTTP/1.1 200 OK
        Content-Type: text/csv
        ...
        1,VOICE_MESSAGES,add,2023-09-01 17:49:03
        ```