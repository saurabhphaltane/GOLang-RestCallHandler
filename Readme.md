Install dependency (Routes)

go get github.com/drone/routes
Run the app

go run app.go
Now, open this URL (http://localhost:3000/profile/foo@gmail.com) in a web browser!



RESTful service implementation in Go Lang to manage personal perferences that can be shared with service providers like travel agencies so that they can provide personalization in their services.

APIs

Create new user profile.
URL: POST /profile

Request

{
    "email": "foo@gmail.com",
    "zip": "94112",
    "country": "U.S.A",
    "profession": "student",
    "favorite_color": "blue",
    "is_smoking": "yes|no",
    "favorite_sport": "hiking",
    "food": {
        "type": "vegetrian|meat_eater|eat_everything",
        "drink_alcohol": "yes|no"
    },
    "music": {
        "spotify_user_id": "wizzler"
    },
    "movie": {
        "tv_shows": ["x", "y", "z"],
        "movies": ["x", "y", "z"]
    },
    "travel": {
        "flight": {
            "seat": "aisle|window"            
        }
    }
}
Response

HTTP/1.1 201 Created
Date: Mon, 29 Feb 2016 19:55:15 GMT
....
Get user profile.
URL: GET /profile/{email}

Response

HTTP/1.1 200 OK
Date: Mon, 29 Feb 2016 19:55:15 GMT
....
{
    "email": "foo@gmail.com",
    "zip": "94112",
    "country": "U.S.A",
    "profession": "student",
    "favorite_color": "blue",
    "is_smoking": "yes",
    "favorite_sport": "hiking",
    "food": {
        "type": "meat_eater",
        "drink_alcohol": "no"
    },
    "music": {
        "spotify_user_id": "wizzler"
    },
    "movie": {
        "tv_shows": ["x", "y", "z"],
        "movies": ["x", "y", "z"]
    },
    "travel": {
        "flight": {
            "seat": "aisle|window"            
        }
    }
}
Update existing user profile.
URL: PUT /profile/{email}

Request

{
    "travel": {
        "flight": {
            "seat": "window"            
        }
    }
}
*** Response ***

HTTP/1.1 204 No Content
Date: Mon, 29 Feb 2016 19:55:15 GMT
....
Delete user profile.
URL: DELETE /profile/{email}

Response

HTTP/1.1 204 No Content
Date: Mon, 29 Feb 2016 19:55:15 GMT
....