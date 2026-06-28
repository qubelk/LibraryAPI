# LibraryAPI

This repository is rewrite of my old repository rest api for library. This demonstrate of my code skills up.

## Differences description

Unlike the previous version, this one uses a PostgreSQL database rather than storing it in memory. It also uses error logging and the gin framework. Overall, the code quality has improved, as has the distribution of responsibilities across packages.

## Differences table

| Future            | Old version                                                         | Current version                                          |
|-------------------|---------------------------------------------------------------------|----------------------------------------------------------|
| Storage           | In RAM. Live while program running.                                 | In PostgreSQL.                                           |
| HTTP              | std http for handlers and chi for routhing.                         | Using gin for everything.                                |
| Package structure | All packages just in root folder. Extra trash in packages.          | Using pkg folder and packages separated by functional.   |
| DTO               | DTO structure are in server package.                                | DTO in its own package. Its simplify thinking for a bit  |
| Validations       | None.                                                               | Have validations of data but simple because no need more |
| Logging           | Not use logging or a output any information.                        | Have logging for a logs/ directory                       |

## Сode limits

This is a base of library API. It lacks user authorization and authentication, sorting search, tests and more. But this is work for CRUD operations in a books.

## For why i writy this

I wanted to rewrite my old code because it is very bad and there is nothing else someday I will rewrite this one too.
