param(
    [string]$IndexName = $( throw 'Index name required' ),
    [string]$MappingFile = $( throw 'Mapping file path required' ),
    [int]$IndexVersion = ${throw 'Index Version required'},
    [boolean]$AddAlias = $False,
    [int]$RemoveAliasIndexVersion = -10000,
    [string]$Uri = 'http://127.0.0.1:9200'
)
$VersionedIndexName = "${IndexName}-v${IndexVersion}"

@{
    'IndexName' = $IndexName;
    'IndexVersion' = $IndexVersion;
    'AddAlias' = $AddAlias;
    'RemoveAliasIndexVersion' = $RemoveAliasIndexVersion;
    'Uri' = $Uri;
    'MappingFile' = $MappingFile;
}

$R = Invoke-WebRequest -Uri "${Uri}/" -Method Get
if ($R.StatusCode -ne 200)
{
    exit 1
}
Write-Host $Uri
Write-Host $R.Content

$R = Invoke-WebRequest -Uri "${Uri}/${VersionedIndexName}?pretty" -Method 'PUT' -InFile "${MappingFile}" -ContentType 'application/json'
if ($R.StatusCode -ne 200)
{
    exit 1
}
Write-Host $R.Content

if ($AddAlias)
{
    $RequestBody = @"
    {
        "actions": [
            {
                "add": {
                    "index": "${VersionedIndexName}",
                    "alias": "${IndexName}"
                }
            }
        ]
    }
"@
    $Data = Invoke-RestMethod -Uri "${Uri}/_aliases?pretty" -ContentType 'application/json' -Method 'Post' -Body $RequestBody
    Write-Host $Data
}

if ($RemoveAliasIndexVersion -ne -10000)
{
    $RemoveVersionedIndexName = "${IndexName}-v${RemoveAliasIndexVersion}"
    $RequestBody = @"
    {
        "actions": [
            {
                "remove": {
                    "index": "${RemoveVersionedIndexName}",
                    "alias": "${IndexName}"
                }
            }
        ]
    }
"@
    $Data = Invoke-RestMethod -Uri "${Uri}/_aliases?pretty" -ContentType 'application/json' -Method 'Post' -Body $RequestBody
    Write-Host $Data
}

Write-Host 'Execute Done.'
exit 0