import smtplib
from email.message import EmailMessage

# Create the email
msg = EmailMessage()
msg.set_content('This is a test email sent to our temporary mailbox service!')
msg['Subject'] = 'Test Email'
msg['From'] = "test@example.com"
msg['To'] = "REPLACE_WITH_YOUR_TEMP_EMAIL@oopsmail.com"  # Replace with your temporary email

# Send the email
try:
    s = smtplib.SMTP('localhost', 1025)
    s.send_message(msg)
    s.quit()
    print("Email sent successfully!")
except Exception as e:
    print(f"Failed to send email: {e}") 