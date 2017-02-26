# goemail
Send email in golang

## Overview
This demo shows how to send emails in golang, which is usually needed to send alarm emails.

## Tips
The codes in the demo are easy to read.

So, what I want to talk about here are some tips when filling in the *sender information* and configuring your *email-setting*.
I will take [gmail](https://mail.google.com/mail) for example.

You can surely fill in *PASSWD* in [mymem.go](https://github.com/seafooler/goemail/blob/master/libofm/mymem.go) with your true passwd. However, it's unsafe.

Instead, you can set an [App password](https://support.google.com/accounts/answer/185833) and replace the *PASSWD* in mymem.go with it:

* Step1: Turn on your [2-step verification](https://support.google.com/accounts/answer/185839)
* Step2: Generate your 16-bit App passwords, and fill in the *PASSWD* with it.

## License 
MIT (see [LICENSE](https://github.com/orcaman/concurrent-map/blob/master/LICENSE) file)