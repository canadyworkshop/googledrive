# googledrive

googledrive is a CLI utility to that can be used working with a Google Drive account from a scripting point of view.
This differs from other projects as it's aim is not for end users but rather as a component in simple automations
where a Google Drive is more appropriate than a Google Storage Bucket.

--service-account-key
upload
download
create directory
delete directory
delete file

file create
-f PathToFile
-n NameOfFile
-e Optional Extension
-d Description
-p A comma separated list of parent folder IDs.
-P A comma separated list of parent paths.

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

https://developers.google.com/drive/api/guides/search-files#examples




