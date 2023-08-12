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

| Short | Long | Type | Default |Details |
|-------|------|-----------------------------------------------------------------------------------------------------|
|q|query| string | |The [query](https://developers.google.com/drive/api/guides/search-files) to apply to the
retrieval. |
|a|all-drives|bool|true|If true files from both My Drive and shared drives will be listed.|
|d|drive-id|string||Limits the results to only files from the drive with the ID provided.|
|f|fields|string||A comma separated list of what fields to return.|
|l|include-labels|bool|false|If true labels will be retrieved.|
|o|order-by|string|A comma separted list of fields to sort the results by.|
|s|page-size|int64|API Default|The maximum number of files to include request to the Google Drive API. By default all
files will be retrieved by making multiple requests. If Paged is set to true then this parameter also controls the
number of files per call.
|p|paged|bool|false|If provided the results will be limited to the page-size. The results will also provide a pageToken
value that can be used with page-token to retrieve the next set of results.|
|n|page-token|string||Requests the next page of results based on the token. This is used in combination with page to
allow manual paging of results.|

-q The Query
-a Query all drives
-d Sets the driveId
-m Set a maximum results per page request
-p Sets the result to be paged.
-f List fields to display.

https://developers.google.com/drive/api/guides/search-files#examples




