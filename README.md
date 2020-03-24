# Form3 Account API Module
Form3 Interview Test - Module Repository

## Developer

Oras Al-Kubaisi

## Objectives
To write a HTTP client for accounts API using Go lang.

## Decisions
I know this is just an interview test but I tried to reflect how I would implement this feature in production environment
so the technical decisions I made was to reflect that and of course to make a supportable and maintainable code.


***1. Create a Go Module in a repo and another application to use the module.***

Reasons:
- Code re-usability. For example if we decided in the future to use one of the endpoints in another application, rather
than re-writing that part, we could just use the module again.
- Separate unit-test from integration test.

***2. Error messages as variables***

Reasons:
- Easy to track all errors at the top of the file.
- Easy to test as we just assert the variable instead of re-writing hard-coded strings. This will save time as well if we decided to change the error message as there will be no need to update unittest.
- Re-usability if the error occurred more than once.
 
***3. Using `Sprintf` for string formatting***

Reason:
- Personal preference as I find them easier to read and maintain instead of chained `+` signs between strings. 

***4. I haven't implemented all business logic to validate `Create` account for all countries***

Reasons:
- Time consuming.
- It will follow the same pattern I have implemented. A new country means a new file inside `accounts` folder and unit test it. Also adding the country code to `switch` statement in `Validate` method.

## Test

```bash
go test -v ./...
```

## Integration test and docker-compose

This will be in the other repository `https://github.com/orasik/form3integration`
