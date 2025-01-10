// internal/tools/mailgun.go
package tools

const MailgunTemplate = `
from restack_ai.function import function, log
from typing import Any
from pydantic import BaseModel

# Add your imports here

# Add this function in services.py file when your function is ready
# to the respective service and workflow config

class %sInput(BaseModel):
    to_address: str
    subject: str = ""
    
    template: str = ""
    message: str = ""
    user_email: str = ""
    data: dict = {}

@function.defn()
async def %s(input: %sInput) -> str:
    """Send a single email using Mailgun API."""
    MAILGUN_API_KEY = os.getenv("MAILGUN_API_KEY")
    MAILGUN_API_URL = os.getenv("MAILGUN_API_URL")
    
    if not MAILGUN_API_KEY or not MAILGUN_API_URL:
        raise ValueError("Missing required environment variables: MAILGUN_API_KEY or MAILGUN_API_URL")

    FROM_EMAIL_ADDRESS = "ADD_YOUR_EMAIL_HERE" 
    
    # Prepare data object
    data = {
        "from": FROM_EMAIL_ADDRESS,
        "to": input.to_address, 
        "subject": input.subject, 
        "h:Reply-To": FROM_EMAIL_ADDRESS,
        "template": input.template,
    }

    variables = {}
    data["h:X-Mailgun-Variables"] = json.dumps(variables) 

    try:
        resp = requests.post(
            MAILGUN_API_URL, 
            auth=("api", MAILGUN_API_KEY),
            data=data
        )
        
        resp.raise_for_status()

        if resp.status_code == 200:
            success_message = f"Successfully sent an email to '{input.to_address}' via Mailgun API."
            log.info(success_message)
            return success_message
        
        error_message = f"Could not send the email, status code: {resp.status_code}, reason: {resp.text}"
        raise Exception(error_message)

    except requests.exceptions.RequestException as e:
        log.error(f"Request Exception: {str(e)}")
        log.error(f"Response Status: {resp.status_code}")
        log.error(f"Response Headers: {dict(resp.headers)}")
        log.error(f"Response Content: {resp.content}")
        
        if resp.status_code == 401:
            log.error("Authentication failed - please check your API key")
        elif resp.status_code == 404:
            log.error("API endpoint not found - please check your API URL")
            
        raise Exception(f"Mailgun API error: {str(e)}")
        
    except Exception as e:
        log.error(f"Failed to send email: {str(e)}")
        raise e
