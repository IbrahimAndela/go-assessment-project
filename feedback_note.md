# Feedback

## Language
Score: 3/5 ==> Meets Expectations

1. Error handling especially in the handlers/articles.go is really bad. These are imminent panics if the application is deployed. Don’t assume incoming data will always Unmarshal without error. Using the struct after a failed unmarshal will lead to a panic. Same goes to a db fetch if no record is found.
2. Candidate forgot to close the request body.
3. Has documentation both README and inline documentation 
4. Has versioned routes 

## Architecture and Design
Score: 3/5 ==> Meets Expectations

1. The candidate has somewhat good design approach on the package level and high level method modularization. However, on the method implementation level, a lot more thought need to be put into separating and modularizing the methods implementation. E.g. UpdateArticle and GetArticle are not modularized.
2. Counter implementation is unnecessarily complex and not optimal. Keep the solution simple and concise.

## Testing and Security
Score: 2/5 ==> Below Expectations

1. The tests covered only verify the endpoint response. This is a catch all approach in integration testing. However, there is a lot of scenarios and use cases that could be missed. A bottom up approach is recommended where the articles.go implementations are modularized and tests done on the particular methods.

# Summary
## Overall Average Rating = 2.7/5
