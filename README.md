# backup2Drive
Go script that uploads a file(s) to a specified Google Drive folder. This file should be compiled. It is built to run on Linux under amd64

## Requirements

  * Google account
  * Drive folder (and its ID)
  * Google Cloud service account
  * credentials file for the above service account

### Go packages
     `go get google.golang.org/api/drive/v3`
     `go get google.golang.org/api/option`

## Drive Folder
Each Google Drive Folder has a unique identifier. To get this ID, in your browser navigate to the folder (if you have not yet created the folder you wish to use, create it now). The folder ID is the last part of the URL in the address bar.

## Create Service Account
backup2Drive uses a service account to access Drive. Service accounts exist in the Google Cloud Platform within a project.
If you do not have a project for this system, create one now.
### Create a project(skip this step if you already have a project)

  * Log into your Google Cloud account
  * Navigate to the console
  * At the top-left, to the right of "Google Cloud Platform", click on the dropdown
  * In the window that opens, click on "NEW PROJECT"
  * Give your project a name
  * Click "CREATE"

### Add a service account
  * From within the Google Cloud Platform, select your project (dropdown, to the right of "Google Cloud Platform")
  * In the left Nav, click "IAM & Admin
  * Click on "Service Accounts"
  * Click "+ Create Service Account"
  * Add a name, ID, and description
  * Don't complete any of the optional info
  * Click "Continue"
  * Click "Done"
  * Click "+ Create Key", then "Create"
  * Select the email address of the service account you just created
  * Select the "Keys" tab
  * In the "Add Key" dropdown, select "Create new key"
  * Click "Create"

When you complete the above, a private/public key pair is generated. Save the file that is generated to your local computer. KEEP UP WITH THE FILE. This is your only opportunity to get the private key.
From the container (Drive or sub-folder) share the Drive folder you are using with the service account
TODO - explain sharing the folder

### Enable Drive API
The GCP project must have the Google Drive API enabled.

  * From within the Google Cloud Platform,  select your project (dropdown, to the right of "Google Cloud Platform")
  * Click on Main Menu->APIs & Services->Library
  * Search for "drive"
  * Click on "Google Drive API"
  * Click "Enable"


## Create the Server-Side configuration
In your editor, create a file named config.json. There is a config.json.example file you can use as a starting point. Copy the folder ID described above. Paste the ID in the config file as the appropriate json value.

```
   {
       "folderID": "Paste folder ID here",
       "files": [
           "/path/to/file1",
           "/path/to/file2",
           "/path/to/etc"
       ]
       // additional configurations
   }
```

Add absolute file paths for each file you wish to upload within the "files" json key.<br><br>

Upload config.json to the /etc/backup2Drive directory on your server. Create the backup2Drive directory, if it doesn't exist. <br>
Upload the file created in "Create Service Account" above, to /etc/backup2Drive/credentials.json on your server.

## Development
TODO - explain dependencies, compiling, etc
