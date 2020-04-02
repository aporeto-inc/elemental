package testmodel

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ListIdentity represents the Identity of the object.
var ListIdentity = elemental.Identity{
	Name:     "list",
	Category: "lists",
	Package:  "todo-list",
	Private:  false,
}

// ListsList represents a list of Lists
type ListsList []*List

// Identity returns the identity of the objects in the list.
func (o ListsList) Identity() elemental.Identity {

	return ListIdentity
}

// Copy returns a pointer to a copy the ListsList.
func (o ListsList) Copy() elemental.Identifiables {

	copy := append(ListsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ListsList.
func (o ListsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ListsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*List))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ListsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ListsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ListsList converted to SparseListsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ListsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseListsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseList)
	}

	return out
}

// Version returns the version of the content.
func (o ListsList) Version() int {

	return 1
}

// List represents the model of a list
type List struct {
	// The identifier.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// This attribute is creation only.
	CreationOnly string `json:"creationOnly" msgpack:"creationOnly" bson:"creationonly" mapstructure:"creationOnly,omitempty"`

	// The date.
	Date time.Time `json:"date" msgpack:"date" bson:"date" mapstructure:"date,omitempty"`

	// The description.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// The name.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// The identifier of the parent of the object.
	ParentID string `json:"parentID" msgpack:"parentID" bson:"parentid" mapstructure:"parentID,omitempty"`

	// The type of the parent of the object.
	ParentType string `json:"parentType" msgpack:"parentType" bson:"parenttype" mapstructure:"parentType,omitempty"`

	// This attribute is readonly.
	ReadOnly string `json:"readOnly" msgpack:"readOnly" bson:"readonly" mapstructure:"readOnly,omitempty"`

	// This attribute is secret.
	Secret string `json:"secret" msgpack:"secret" bson:"secret" mapstructure:"secret,omitempty"`

	// this is a slice.
	Slice []string `json:"slice" msgpack:"slice" bson:"slice" mapstructure:"slice,omitempty"`

	// This attribute is not exposed.
	Unexposed string `json:"-" msgpack:"-" bson:"unexposed" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewList returns a new *List
func NewList() *List {

	return &List{
		ModelVersion: 1,
		Slice:        []string{},
	}
}

// Identity returns the Identity of the object.
func (o *List) Identity() elemental.Identity {

	return ListIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *List) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *List) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *List) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesList{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.CreationOnly = o.CreationOnly
	s.Date = o.Date
	s.Description = o.Description
	s.Name = o.Name
	s.ParentID = o.ParentID
	s.ParentType = o.ParentType
	s.ReadOnly = o.ReadOnly
	s.Secret = o.Secret
	s.Slice = o.Slice
	s.Unexposed = o.Unexposed

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *List) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesList{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.CreationOnly = s.CreationOnly
	o.Date = s.Date
	o.Description = s.Description
	o.Name = s.Name
	o.ParentID = s.ParentID
	o.ParentType = s.ParentType
	o.ReadOnly = s.ReadOnly
	o.Secret = s.Secret
	o.Slice = s.Slice
	o.Unexposed = s.Unexposed

	return nil
}

// Version returns the hardcoded version of the model.
func (o *List) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *List) BleveType() string {

	return "list"
}

// DefaultOrder returns the list of default ordering fields.
func (o *List) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *List) Doc() string {

	return `Represent a a list of task to do.`
}

func (o *List) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetName returns the Name of the receiver.
func (o *List) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *List) SetName(name string) {

	o.Name = name
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *List) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseList{
			ID:           &o.ID,
			CreationOnly: &o.CreationOnly,
			Date:         &o.Date,
			Description:  &o.Description,
			Name:         &o.Name,
			ParentID:     &o.ParentID,
			ParentType:   &o.ParentType,
			ReadOnly:     &o.ReadOnly,
			Secret:       &o.Secret,
			Slice:        &o.Slice,
			Unexposed:    &o.Unexposed,
		}
	}

	sp := &SparseList{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "creationOnly":
			sp.CreationOnly = &(o.CreationOnly)
		case "date":
			sp.Date = &(o.Date)
		case "description":
			sp.Description = &(o.Description)
		case "name":
			sp.Name = &(o.Name)
		case "parentID":
			sp.ParentID = &(o.ParentID)
		case "parentType":
			sp.ParentType = &(o.ParentType)
		case "readOnly":
			sp.ReadOnly = &(o.ReadOnly)
		case "secret":
			sp.Secret = &(o.Secret)
		case "slice":
			sp.Slice = &(o.Slice)
		case "unexposed":
			sp.Unexposed = &(o.Unexposed)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseList to the object.
func (o *List) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseList)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.CreationOnly != nil {
		o.CreationOnly = *so.CreationOnly
	}
	if so.Date != nil {
		o.Date = *so.Date
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.ParentID != nil {
		o.ParentID = *so.ParentID
	}
	if so.ParentType != nil {
		o.ParentType = *so.ParentType
	}
	if so.ReadOnly != nil {
		o.ReadOnly = *so.ReadOnly
	}
	if so.Secret != nil {
		o.Secret = *so.Secret
	}
	if so.Slice != nil {
		o.Slice = *so.Slice
	}
	if so.Unexposed != nil {
		o.Unexposed = *so.Unexposed
	}
}

// DeepCopy returns a deep copy if the List.
func (o *List) DeepCopy() *List {

	if o == nil {
		return nil
	}

	out := &List{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *List.
func (o *List) DeepCopyInto(out *List) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy List: %s", err))
	}

	*out = *target.(*List)
}

// Validate valides the current information stored into the structure.
func (o *List) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*List) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ListAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ListLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*List) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ListAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *List) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "creationOnly":
		return o.CreationOnly
	case "date":
		return o.Date
	case "description":
		return o.Description
	case "name":
		return o.Name
	case "parentID":
		return o.ParentID
	case "parentType":
		return o.ParentType
	case "readOnly":
		return o.ReadOnly
	case "secret":
		return o.Secret
	case "slice":
		return o.Slice
	case "unexposed":
		return o.Unexposed
	}

	return nil
}

// ListAttributesMap represents the map of attribute for List.
var ListAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `The identifier.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"CreationOnly": {
		AllowedChoices: []string{},
		ConvertedName:  "CreationOnly",
		CreationOnly:   true,
		Description:    `This attribute is creation only.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "creationOnly",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Date": {
		AllowedChoices: []string{},
		ConvertedName:  "Date",
		Description:    `The date.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "date",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"Description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `The description.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `The name.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ParentID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentID",
		Description:    `The identifier of the parent of the object.`,
		Exposed:        true,
		Filterable:     true,
		ForeignKey:     true,
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ParentType": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentType",
		Description:    `The type of the parent of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "parentType",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ReadOnly": {
		AllowedChoices: []string{},
		ConvertedName:  "ReadOnly",
		Description:    `This attribute is readonly.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "readOnly",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Secret": {
		AllowedChoices: []string{},
		ConvertedName:  "Secret",
		Description:    `This attribute is secret.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "secret",
		Orderable:      true,
		Secret:         true,
		Stored:         true,
		Type:           "string",
	},
	"Slice": {
		AllowedChoices: []string{},
		ConvertedName:  "Slice",
		Description:    `this is a slice.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "slice",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Unexposed": {
		AllowedChoices: []string{},
		ConvertedName:  "Unexposed",
		Description:    `This attribute is not exposed.`,
		Filterable:     true,
		Name:           "unexposed",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}

// ListLowerCaseAttributesMap represents the map of attribute for List.
var ListLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `The identifier.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"creationonly": {
		AllowedChoices: []string{},
		ConvertedName:  "CreationOnly",
		CreationOnly:   true,
		Description:    `This attribute is creation only.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "creationOnly",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"date": {
		AllowedChoices: []string{},
		ConvertedName:  "Date",
		Description:    `The date.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "date",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `The description.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `The name.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"parentid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentID",
		Description:    `The identifier of the parent of the object.`,
		Exposed:        true,
		Filterable:     true,
		ForeignKey:     true,
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"parenttype": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentType",
		Description:    `The type of the parent of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "parentType",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"readonly": {
		AllowedChoices: []string{},
		ConvertedName:  "ReadOnly",
		Description:    `This attribute is readonly.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "readOnly",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"secret": {
		AllowedChoices: []string{},
		ConvertedName:  "Secret",
		Description:    `This attribute is secret.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "secret",
		Orderable:      true,
		Secret:         true,
		Stored:         true,
		Type:           "string",
	},
	"slice": {
		AllowedChoices: []string{},
		ConvertedName:  "Slice",
		Description:    `this is a slice.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "slice",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"unexposed": {
		AllowedChoices: []string{},
		ConvertedName:  "Unexposed",
		Description:    `This attribute is not exposed.`,
		Filterable:     true,
		Name:           "unexposed",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}

// SparseListsList represents a list of SparseLists
type SparseListsList []*SparseList

// Identity returns the identity of the objects in the list.
func (o SparseListsList) Identity() elemental.Identity {

	return ListIdentity
}

// Copy returns a pointer to a copy the SparseListsList.
func (o SparseListsList) Copy() elemental.Identifiables {

	copy := append(SparseListsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseListsList.
func (o SparseListsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseListsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseList))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseListsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseListsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseListsList converted to ListsList.
func (o SparseListsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseListsList) Version() int {

	return 1
}

// SparseList represents the sparse version of a list.
type SparseList struct {
	// The identifier.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// This attribute is creation only.
	CreationOnly *string `json:"creationOnly,omitempty" msgpack:"creationOnly,omitempty" bson:"creationonly,omitempty" mapstructure:"creationOnly,omitempty"`

	// The date.
	Date *time.Time `json:"date,omitempty" msgpack:"date,omitempty" bson:"date,omitempty" mapstructure:"date,omitempty"`

	// The description.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// The name.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// The identifier of the parent of the object.
	ParentID *string `json:"parentID,omitempty" msgpack:"parentID,omitempty" bson:"parentid,omitempty" mapstructure:"parentID,omitempty"`

	// The type of the parent of the object.
	ParentType *string `json:"parentType,omitempty" msgpack:"parentType,omitempty" bson:"parenttype,omitempty" mapstructure:"parentType,omitempty"`

	// This attribute is readonly.
	ReadOnly *string `json:"readOnly,omitempty" msgpack:"readOnly,omitempty" bson:"readonly,omitempty" mapstructure:"readOnly,omitempty"`

	// This attribute is secret.
	Secret *string `json:"secret,omitempty" msgpack:"secret,omitempty" bson:"secret,omitempty" mapstructure:"secret,omitempty"`

	// this is a slice.
	Slice *[]string `json:"slice,omitempty" msgpack:"slice,omitempty" bson:"slice,omitempty" mapstructure:"slice,omitempty"`

	// This attribute is not exposed.
	Unexposed *string `json:"-" msgpack:"-" bson:"unexposed,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseList returns a new  SparseList.
func NewSparseList() *SparseList {
	return &SparseList{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseList) Identity() elemental.Identity {

	return ListIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseList) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseList) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseList) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseList{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.CreationOnly != nil {
		s.CreationOnly = o.CreationOnly
	}
	if o.Date != nil {
		s.Date = o.Date
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.ParentID != nil {
		s.ParentID = o.ParentID
	}
	if o.ParentType != nil {
		s.ParentType = o.ParentType
	}
	if o.ReadOnly != nil {
		s.ReadOnly = o.ReadOnly
	}
	if o.Secret != nil {
		s.Secret = o.Secret
	}
	if o.Slice != nil {
		s.Slice = o.Slice
	}
	if o.Unexposed != nil {
		s.Unexposed = o.Unexposed
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseList) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseList{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.CreationOnly != nil {
		o.CreationOnly = s.CreationOnly
	}
	if s.Date != nil {
		o.Date = s.Date
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.ParentID != nil {
		o.ParentID = s.ParentID
	}
	if s.ParentType != nil {
		o.ParentType = s.ParentType
	}
	if s.ReadOnly != nil {
		o.ReadOnly = s.ReadOnly
	}
	if s.Secret != nil {
		o.Secret = s.Secret
	}
	if s.Slice != nil {
		o.Slice = s.Slice
	}
	if s.Unexposed != nil {
		o.Unexposed = s.Unexposed
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseList) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseList) ToPlain() elemental.PlainIdentifiable {

	out := NewList()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.CreationOnly != nil {
		out.CreationOnly = *o.CreationOnly
	}
	if o.Date != nil {
		out.Date = *o.Date
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.ParentID != nil {
		out.ParentID = *o.ParentID
	}
	if o.ParentType != nil {
		out.ParentType = *o.ParentType
	}
	if o.ReadOnly != nil {
		out.ReadOnly = *o.ReadOnly
	}
	if o.Secret != nil {
		out.Secret = *o.Secret
	}
	if o.Slice != nil {
		out.Slice = *o.Slice
	}
	if o.Unexposed != nil {
		out.Unexposed = *o.Unexposed
	}

	return out
}

// GetName returns the Name of the receiver.
func (o *SparseList) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseList) SetName(name string) {

	o.Name = &name
}

// DeepCopy returns a deep copy if the SparseList.
func (o *SparseList) DeepCopy() *SparseList {

	if o == nil {
		return nil
	}

	out := &SparseList{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseList.
func (o *SparseList) DeepCopyInto(out *SparseList) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseList: %s", err))
	}

	*out = *target.(*SparseList)
}

type mongoAttributesList struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	CreationOnly string        `bson:"creationonly"`
	Date         time.Time     `bson:"date"`
	Description  string        `bson:"description"`
	Name         string        `bson:"name"`
	ParentID     string        `bson:"parentid"`
	ParentType   string        `bson:"parenttype"`
	ReadOnly     string        `bson:"readonly"`
	Secret       string        `bson:"secret"`
	Slice        []string      `bson:"slice"`
	Unexposed    string        `bson:"unexposed"`
}
type mongoAttributesSparseList struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	CreationOnly *string       `bson:"creationonly,omitempty"`
	Date         *time.Time    `bson:"date,omitempty"`
	Description  *string       `bson:"description,omitempty"`
	Name         *string       `bson:"name,omitempty"`
	ParentID     *string       `bson:"parentid,omitempty"`
	ParentType   *string       `bson:"parenttype,omitempty"`
	ReadOnly     *string       `bson:"readonly,omitempty"`
	Secret       *string       `bson:"secret,omitempty"`
	Slice        *[]string     `bson:"slice,omitempty"`
	Unexposed    *string       `bson:"unexposed,omitempty"`
}
