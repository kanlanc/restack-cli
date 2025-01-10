// internal/templates/function.go
package templates

const FunctionTemplate = `
from restack_ai.function import function, log
from typing import Any
from pydantic import BaseModel

# Add your imports here

# Add this function in services.py file when your function is ready
# to the respective service and workflow config

class %sInput(BaseModel):
    # Add your input fields here
    pass

@function.defn()
async def %s(input: %sInput):
    try:
        # Your function logic here
        pass
    except Exception as error:
        log.error("%s function failed", error=error)
        raise error
`
