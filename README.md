## CDR-Analytics
CDR reporting for Rosita


This is a web-based tool for analyzing Call Detail Records (CDRs) stored in a MySQL database. The tool is built using the Go programming language, the Gin web framework, the Gorm ORM, and various other open-source libraries.

#### Features

    Filter CDRs based on various criteria such as date range, call type, call status, and call direction.
    Generate various types of reports based on the filtered CDRs, including charts and graphs.
    Export filtered CDRs to various formats such as CSV or Excel.
    Monitor real-time call status and display metrics such as the number of active calls, waiting calls, and answered calls.
    Manage user accounts, roles, and permissions.
    Display metrics and charts on a user-friendly dashboard.

#### Installation

    Clone this repository to your local machine using git clone https://github.com/YOUR-USERNAME/cdr-reporting-tool.git
    Change directory to the cloned repository using cd cdr-reporting-tool
    Install the required dependencies using go mod download

#### Configuration

The tool can be configured using a config.json file, which should be placed in the root directory of the repository. Here is a sample configuration file:
```
json

{
  "database": {
    "driver": "mysql",
    "host": "localhost",
    "port": 3306,
    "username": "root",
    "password": "",
    "database": "cdr_db",
    "timezone": "UTC"
  },
  "server": {
    "host": "localhost",
    "port": 8080
  },
  "jwt": {
    "secret": "my-secret-key",
    "expiration": 86400
  }
}
```

Usage

    Start the web server using go run main.go.
    Open a web browser and navigate to http://localhost:8080.

#### API Reference

    Authentication API: /auth
    CDR retrieval API: /cdrs
    CDR filtering API: /cdrs?start_date=2022-01-01&end_date=2022-01-31&call_type=inbound
    Reporting API: /reports?type=bar&data=active_calls
    Export API: /export?type=csv
    Real-time monitoring API: /monitor
    Configuration API: /config
    User management API: /users
    Dashboard API: /dashboard

##### CDR Retrieval API

The CDR retrieval API allows you to retrieve call detail records (CDRs) from the database. You can retrieve all CDRs or filter them based on various criteria such as date range, call type, call status, and call direction.
Endpoint

bash

GET /cdrs

Query Parameters

    start_date: The start date of the date range to filter the CDRs by (e.g., 2022-01-01).
    end_date: The end date of the date range to filter the CDRs by (e.g., 2022-01-31).
    call_type: The type of call to filter the CDRs by (e.g., inbound or outbound).
    call_status: The status of the call to filter the CDRs by (e.g., answered or unanswered).
    call_direction: The direction of the call to filter the CDRs by (e.g., incoming or outgoing).

Response

The API returns a JSON object containing an array of CDRs that match the specified filter criteria.

##### CDR Filtering API

The CDR filtering API allows you to filter CDRs based on various criteria such as date range, call type, call status, and call direction. You can use this API to retrieve CDRs that match the specified filter criteria.
Endpoint

bash

GET /cdrs?start_date=2022-01-01&end_date=2022-01-31&call_type=inbound

Query Parameters

    start_date: The start date of the date range to filter the CDRs by (e.g., 2022-01-01).
    end_date: The end date of the date range to filter the CDRs by (e.g., 2022-01-31).
    call_type: The type of call to filter the CDRs by (e.g., inbound or outbound).
    call_status: The status of the call to filter the CDRs by (e.g., answered or unanswered).
    call_direction: The direction of the call to filter the CDRs by (e.g., incoming or outgoing).

Response

The API returns a JSON object containing an array of CDRs that match the specified filter criteria.

##### Reporting API

The reporting API allows you to generate various types of reports based on the filtered CDRs. You can generate charts and graphs that visualize the call data.
Endpoint

bash

GET /reports?type=bar&data=active_calls

Query Parameters

    type: The type of chart or graph to generate (e.g., bar or pie).
    data: The type of data to display on the chart or graph (e.g., active_calls or answered_calls).

Response

The API returns a JSON object containing the data necessary to generate the specified chart or graph. The front-end can use this data to render the chart or graph.

#### Contributing

Contributions are welcome! Please open an issue or pull request if you have any suggestions or bug reports.
License

This project is licensed under the MIT License - see the LICENSE file for details.
