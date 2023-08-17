# googledrive

googledrive is a CLI utility that's aim is to be used with scripting. This differs from other Google Drive CLI
utilities that primarily focus on providing file syncing.


# file list

| Long           | Type   | Default                                                 | Details                                                                                                                                                                                                                             |
|----------------|--------|---------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| query          | string |                                                         | The [query](https://developers.google.com/drive/api/guides/search-files) to apply to the                                                                                                                                            |
| all-drives     | bool   | true                                                    | If true files from both My Drive and shared drives will be listed.                                                                                                                                                                  |
| drive-id       | string |                                                         | Limits the results to only files from the drive with the ID provided.                                                                                                                                                               |
| fields         | string |                                                         | A comma separated list of what fields to return.                                                                                                                                                                                    |
| include-labels | string |                                                         | If true labels will be retrieved.                                                                                                                                                                                                   |
| order-by       | string | A comma separted list of fields to sort the results by. |                                                                                                                                                                                                                                     |
| page-size      | int64  | API Default                                             | The maximum number of files to include request to the Google Drive API. By default all files will be retrieved by making multiple requests. If Paged is set to true then this parameter also controls the number of files per call. |
| paged          | bool   | false                                                   | If provided the results will be limited to the page-size. The results will also provide a pageToken value that can be used with page-token to retrieve the next set of results.                                                     |
| page-token     | string |                                                         | Requests the next page of results based on the token. This is used in combination with page to allow manual paging of results.                                                                                                      |                                                                                                   |
| list           | bool   | false                                                   | Displays the files in long format.                                                                                                                                                                                                  |
| out            | string | std                                                     | Denotes the format the data should be returned in. [std                                                                                                                                                                             |json]|

# file create

| Long        | Type   | Default                       | Required | Details                                                                                                |
|-------------|--------|-------------------------------|----------|--------------------------------------------------------------------------------------------------------|
| file        | string |                               | true     | The path to the file to upload.                                                                        |
| name        | string | The name of the file in file. | false    | An alternative name for the file. If empty the name of the uploading file will be used.                |
| parents     | string | .                             | false    | The ID of the parent container for the file. If not provided the root of the users drive will be used. |
| description | string |                               | false    | Sets the description field of the file.                                                                |


