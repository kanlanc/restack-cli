// internal/templates/workflow.go
package templates

const WorkflowTemplate = `
import asyncio
from datetime import timedelta
from typing import Any
from pydantic import BaseModel
from restack_ai.workflow import workflow, log, workflow_info, import_functions


with import_functions():
    # from src.functions.<your_function_filename> import <your_function_name>

class Input(BaseModel):
    # Add your input fields here
    pass



@workflow.defn()
class %sWorkflow:
    @workflow.run
    async def run(self, input:Input):
        # Your workflow logic here

		# Call your function as steps like this


		# NOTE: All of function inputs should be passed in a single object

		# input = {
		#	"<input_name>": "<input_value>",
		#}


		# results = await workflow.step(
        #     <function_name>,
        #     input,
        #     start_to_close_timeout=timedelta(minutes=2)
        # )

        return {
            "result": "success"
        }
`
