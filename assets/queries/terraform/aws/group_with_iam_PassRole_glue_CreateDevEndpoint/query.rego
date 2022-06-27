package Cx

import data.generic.common as common_lib

CxPolicy[result] {

	# get a AWS IAM group
	input.document[i].resource.aws_iam_group[targetGroup]

    common_lib.group_unrecommended_permission_policy_scenarios(targetGroup, "glue:CreateDevEndpoint")
    common_lib.group_unrecommended_permission_policy_scenarios(targetGroup, "iam:PassRole")


	result := {
		"documentId": input.document[i].id,
		"searchKey": sprintf("aws_iam_group[%s]", [targetGroup]),
		"issueType": "IncorrectValue",
        "keyExpectedValue": sprintf("group %s is not associated with a policy that has Action set to 'glue:CreateDevEndpoint' and 'iam:PassRole' and Resource set to '*'", [targetGroup]),
		"keyActualValue": sprintf("group %s is associated with a policy that has Action set to 'glue:CreateDevEndpoint' and 'iam:PassRole' and Resource set to '*'", [targetGroup]),
        "searchLine": common_lib.build_search_line(["resource", "aws_iam_group", targetGroup], []),
	}
}
