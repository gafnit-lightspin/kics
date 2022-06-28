package Cx

import data.generic.terraform as tf_lib

CxPolicy[result] {
	resource := input.document[i].resource.azurerm_mysql_server[var0]
	not resource.ssl_enforcement_enabled

	result := {
		"documentId": input.document[i].id,
		"resourceType": "azurerm_mssql_server",
		"resourceName": tf_lib.get_resource_name(resource, var0),
		"searchKey": sprintf("azurerm_mysql_server[%s].ssl_enforcement_enabled", [var0]),
		"issueType": "IncorrectValue",
		"keyExpectedValue": sprintf("'azurerm_mysql_server.%s.ssl_enforcement_enabled' is equal 'true'", [var0]),
		"keyActualValue": sprintf("'azurerm_mysql_server.%s.ssl_enforcement_enabled' is equal 'false'", [var0]),
	}
}
