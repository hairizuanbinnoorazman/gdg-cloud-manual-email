# Email tool

This is an initial prototype of a email tool that is to help organizations to send emails to members. 

Note: DO NOT USE THIS YET AS THERE ARE STILL EXPECTED ISSUES

# Sample commands

The default files that are required here is:

- `email_list_file`
- `email_template.html`

Note, the cli allows users to be able to switch which file to be used

```bash
export SENDGRID_KEY=sample.key
tool email --dryRun=false --subject="Default Subject"
```