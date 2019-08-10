param(
    [string]$IndexName = $( throw 'Index name required' ),
    [int]$IndexVersion = ${throw 'Index Version required'},
    [string]$Uri = 'http://127.0.0.1:9200'
)
$VersionedIndexName = "${IndexName}-v${IndexVersion}"

@{
    'IndexName' = $IndexName;
    'IndexVersion' = $IndexVersion;
    'Uri' = $Uri;
}

$R = Invoke-WebRequest -Uri "${Uri}/" -Method Get
if ($R.StatusCode -ne 200)
{
    exit 1
}
Write-Host $Uri
Write-Host $R.Content

$R = Invoke-WebRequest -Uri "${Uri}/${VersionedIndexName}?pretty" -Method 'Delete' -ContentType 'application/json'
if ($R.StatusCode -ne 200)
{
    exit 1
}
Write-Host $R.Content

Write-Host 'Execute Done.'
exit 0