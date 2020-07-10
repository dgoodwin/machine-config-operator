package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_BrokerTemplateInstance = map[string]string{
	"":     "BrokerTemplateInstance holds the service broker-related state associated with a TemplateInstance.  BrokerTemplateInstance is part of an experimental API.",
	"spec": "spec describes the state of this BrokerTemplateInstance.",
}

func (BrokerTemplateInstance) SwaggerDoc() map[string]string {
	return map_BrokerTemplateInstance
}

var map_BrokerTemplateInstanceList = map[string]string{
	"":      "BrokerTemplateInstanceList is a list of BrokerTemplateInstance objects.",
	"items": "items is a list of BrokerTemplateInstances",
}

func (BrokerTemplateInstanceList) SwaggerDoc() map[string]string {
	return map_BrokerTemplateInstanceList
}

var map_BrokerTemplateInstanceSpec = map[string]string{
	"":                 "BrokerTemplateInstanceSpec describes the state of a BrokerTemplateInstance.",
	"templateInstance": "templateinstance is a reference to a TemplateInstance object residing in a namespace.",
	"secret":           "secret is a reference to a Secret object residing in a namespace, containing the necessary template parameters.",
	"bindingIDs":       "bindingids is a list of 'binding_id's provided during successive bind calls to the template service broker.",
}

func (BrokerTemplateInstanceSpec) SwaggerDoc() map[string]string {
	return map_BrokerTemplateInstanceSpec
}

var map_Parameter = map[string]string{
	"":            "Parameter defines a name/value variable that is to be processed during the Template to Config transformation.",
	"name":        "Name must be set and it can be referenced in Template Items using ${PARAMETER_NAME}. Required.",
	"displayName": "Optional: The name that will show in UI instead of parameter 'Name'",
	"description": "Description of a parameter. Optional.",
	"value":       "Value holds the Parameter data. If specified, the generator will be ignored. The value replaces all occurrences of the Parameter ${Name} expression during the Template to Config transformation. Optional.",
	"generate":    "generate specifies the generator to be used to generate random string from an input value specified by From field. The result string is stored into Value field. If empty, no generator is being used, leaving the result Value untouched. Optional.\n\nThe only supported generator is \"expression\", which accepts a \"from\" value in the form of a simple regular expression containing the range expression \"[a-zA-Z0-9]\", and the length expression \"a{length}\".\n\nExamples:\n\nfrom             | value",
	"from":        "From is an input value for the generator. Optional.",
	"required":    "Optional: Indicates the parameter must have a value.  Defaults to false.",
}

func (Parameter) SwaggerDoc() map[string]string {
	return map_Parameter
}

var map_Template = map[string]string{
	"":           "Template contains the inputs needed to produce a Config.",
	"message":    "message is an optional instructional message that will be displayed when this template is instantiated. This field should inform the user how to utilize the newly created resources. Parameter substitution will be performed on the message before being displayed so that generated credentials and other parameters can be included in the output.",
	"objects":    "objects is an array of resources to include in this template. If a namespace value is hardcoded in the object, it will be removed during template instantiation, however if the namespace value is, or contains, a ${PARAMETER_REFERENCE}, the resolved value after parameter substitution will be respected and the object will be created in that namespace.",
	"parameters": "parameters is an optional array of Parameters used during the Template to Config transformation.",
	"labels":     "labels is a optional set of labels that are applied to every object during the Template to Config transformation.",
}

func (Template) SwaggerDoc() map[string]string {
	return map_Template
}

var map_TemplateInstance = map[string]string{
	"":       "TemplateInstance requests and records the instantiation of a Template. TemplateInstance is part of an experimental API.",
	"spec":   "spec describes the desired state of this TemplateInstance.",
	"status": "status describes the current state of this TemplateInstance.",
}

func (TemplateInstance) SwaggerDoc() map[string]string {
	return map_TemplateInstance
}

var map_TemplateInstanceCondition = map[string]string{
	"":                   "TemplateInstanceCondition contains condition information for a TemplateInstance.",
	"type":               "Type of the condition, currently Ready or InstantiateFailure.",
	"status":             "Status of the condition, one of True, False or Unknown.",
	"lastTransitionTime": "LastTransitionTime is the last time a condition status transitioned from one state to another.",
	"reason":             "Reason is a brief machine readable explanation for the condition's last transition.",
	"message":            "Message is a human readable description of the details of the last transition, complementing reason.",
}

func (TemplateInstanceCondition) SwaggerDoc() map[string]string {
	return map_TemplateInstanceCondition
}

var map_TemplateInstanceList = map[string]string{
	"":      "TemplateInstanceList is a list of TemplateInstance objects.",
	"items": "items is a list of Templateinstances",
}

func (TemplateInstanceList) SwaggerDoc() map[string]string {
	return map_TemplateInstanceList
}

var map_TemplateInstanceObject = map[string]string{
	"":    "TemplateInstanceObject references an object created by a TemplateInstance.",
	"ref": "ref is a reference to the created object.  When used under .spec, only name and namespace are used; these can contain references to parameters which will be substituted following the usual rules.",
}

func (TemplateInstanceObject) SwaggerDoc() map[string]string {
	return map_TemplateInstanceObject
}

var map_TemplateInstanceRequester = map[string]string{
	"":         "TemplateInstanceRequester holds the identity of an agent requesting a template instantiation.",
	"username": "username uniquely identifies this user among all active users.",
	"uid":      "uid is a unique value that identifies this user across time; if this user is deleted and another user by the same name is added, they will have different UIDs.",
	"groups":   "groups represent the groups this user is a part of.",
	"extra":    "extra holds additional information provided by the authenticator.",
}

func (TemplateInstanceRequester) SwaggerDoc() map[string]string {
	return map_TemplateInstanceRequester
}

var map_TemplateInstanceSpec = map[string]string{
	"":          "TemplateInstanceSpec describes the desired state of a TemplateInstance.",
	"template":  "template is a full copy of the template for instantiation.",
	"secret":    "secret is a reference to a Secret object containing the necessary template parameters.",
	"requester": "requester holds the identity of the agent requesting the template instantiation.",
}

func (TemplateInstanceSpec) SwaggerDoc() map[string]string {
	return map_TemplateInstanceSpec
}

var map_TemplateInstanceStatus = map[string]string{
	"":           "TemplateInstanceStatus describes the current state of a TemplateInstance.",
	"conditions": "conditions represent the latest available observations of a TemplateInstance's current state.",
	"objects":    "Objects references the objects created by the TemplateInstance.",
}

func (TemplateInstanceStatus) SwaggerDoc() map[string]string {
	return map_TemplateInstanceStatus
}

var map_TemplateList = map[string]string{
	"":      "TemplateList is a list of Template objects.",
	"items": "Items is a list of templates",
}

func (TemplateList) SwaggerDoc() map[string]string {
	return map_TemplateList
}

// AUTO-GENERATED FUNCTIONS END HERE
