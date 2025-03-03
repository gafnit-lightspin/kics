package Cx

import data.generic.common as common_lib

CxPolicy[result] {

	# get a AWS IAM group
	input.document[i].resource.aws_iam_group[targetGroup]

    common_lib.group_unrecommended_permission_policy_scenarios(targetGroup, "iam:CreateLoginProfile")


	result := {
		"documentId": input.document[i].id,
		"searchKey": sprintf("aws_iam_group[%s]", [targetGroup]),
		"issueType": "IncorrectValue",
        "keyExpectedValue": sprintf("group %s is not associated with a policy that has Action set to 'iam:CreateLoginProfile' and Resource set to '*'", [targetGroup]),
		"keyActualValue": sprintf("group %s is associated with a policy that has Action set to 'iam:CreateLoginProfile' and Resource set to '*'", [targetGroup]),
        "searchLine": common_lib.build_search_line(["resource", "aws_iam_group", targetGroup], []),
	}
}
