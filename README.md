# RideWise: Ride-Hailing App

RideWise is a gRPC microservices ride-hailing Go backend that connects passengers with nearby drivers for convenient and efficient transportation.  Users can easily request rides by specifying their pickup and drop-off locations, with the application employing sophisticated matching algorithms to pair them with available drivers based on factors like proximity, destination, and driver ratings. Real-time tracking functionalities allow passengers to monitor their driver's location and estimated time of arrival, ensuring a seamless and transparent experience. Cashless transactions are facilitated through integrated payment gateways. Administrative services provides oversight, allowing administrators to manage users, drivers, rides, and resolve disputes, while an analytics service offers insights into key metrics for performance optimization. 

Ridewise implements the [API gateway microservices pattern](https://microservices.io/patterns/apigateway.html). The Gateway service serves as the single entry point for all clients into the backend. The Gateway service handles requests in one of two ways: Some requests are simply proxied/routed to the appropriate service, while other requests are handled by fanning out to multiple services and composing a singular response. 

# Service Decomposition: 
A granular decomposition of services:

1. Authentication Service:
    - Handles user authentication and token generation/validation.
    - Utilizes JWT for authentication mechanisms.
    - Implements 2FA via OTP.
    - Exposes endpoint for user login and token management.

2. Rider Management Services:
    - Rider Service:
        - Manages passenger profiles, including CRUD operations on user accounts, and profile updates.
        - Exposes endpoints for passenger registration, and passenger profile management.
    - Authorization Service:
        - Handles access control and authorization policies, including role-based access control (RBAC) and permissions management.
        - Exposes endpoints for managing authorization rules and user permissions.

3. Ride Services:
    - Ride Request Service:
        - Handles ride request management, including matching algorithms and fare estimation.
        - Exposes endpoints for ride requests, ride matching, fare estimation, and ride history.
    - Ride Tracking Service:
        - Manages real-time tracking of driver locations, ETA calculation, and historical tracking data storage.
        - Exposes endpoints for real-time driver tracking and ETA calculation.
    - Payment Service:
        - Integrates with payment gateways for cashless transactions between riders and drivers.
        - Exposes endpoints for payment processing, including payment authorization, capture, and refunds.

4. Driver Management Services:
    - Driver Registration Service:
        - Handles driver registration and onboarding processes.
        - Exposes endpoints for driver registration and onboarding.
    - Driver Profile Service:
        - Manages driver profiles, including CRUD operations on driver accounts, profile updates, and availability management.
        - Exposes endpoints for driver profile management.
    - Vehicle Service:
        - Manages information about driver vehicles, including registration, insurance, and vehicle characteristics.
        - Exposes endpoints for managing driver vehicles.

5. Notification Services:
    - Push Notification Service:
        - Sends push notifications to users for ride updates, payment confirmations, and other relevant events.
        - Exposes endpoints for sending push notifications.
    - Email Notification Service:
        - Sends email notifications to users for account-related activities, such as password resets and account verification.
        - Exposes endpoints for sending email notifications.
    - SMS Notification Service:
        - Sends SMS notifications to users for urgent alerts or notifications where SMS is preferred over other channels.
        - Exposes endpoints for sending SMS notifications.

6. Admin Services:
    - User Administration Service:
        - Handles administrative tasks related to user management, such as user creation, deletion, and role assignment.
        - Exposes endpoints for managing users and roles.
    - Ride Administration Service:
        - Manages administrative tasks related to ride management, including ride history, dispute resolution, and analytics access.
        - Exposes endpoints for managing rides and analytics.
    - System Configuration Service:
        - Manages system-wide configurations and settings, such as feature toggles, logging levels, and environment variables.
        - Exposes endpoints for managing system configurations.

7. Analytics Service:
    - Collects and analyzes data from various microservices to track key metrics like total rides, active users, revenue, etc.
    - Generates reports and visualizations for insights into the performance of the ride-hailing app.
    - Exposes endpoints for querying analytics data and generating reports.
