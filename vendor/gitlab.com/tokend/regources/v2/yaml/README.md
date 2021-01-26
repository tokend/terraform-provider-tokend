# YAML resource specifications
All resources must be specified in YAML to allow code generation.

There are two types of entities: 

1.  Resources in `resources` directory;
2.  Non-resource inner objects in `inner` directory.

Inner objects are object used in resource attributes or in other inner objects such as `Fee`, `XdrEnumValue`, etc.

## Describing a resource
Resource must be described as a `.yaml` file in `resources` directory.

Resource specification has the following structure:

|Name|Description|
|----|-----------|
|name|Human-readable resource name i.e. `Asset`, `Balance`, etc|
|key|Resource key i.e. `assets`, `balances`, etc|
| base |Optional field, describes resources inheritance. Contains a base resource key for given resource|
|attributes|Optional field, list of resource attributes|
|relations|Optional field, list of resource relationships|

### Resource attribute
Resource attribute has the following properties:

|Name|Description|
|----|-----------|
| name |JSON name of the attribute i.e. `pre_issuance_asset_signer`|
| type |Attribute type, see allowed values below|
| optional |Optional field, `true` indicates that this attribute may be missing from the resource|
| is_collection |Optional field, 'true' indicates that this attribute is a list of given type|

Attribute type is a language-independent value. For primitives use the following types:

* `Int32`
* `Int64`
* `UInt32`
* `UInt64`
* `String`
* `Amount` – for string amount representation
* `Date` – for string date representation
* `Object` – for arbitrary non-typed data

Some attributes have a complex structure, for example fees, XDR enum values, XDR bitmasks, etc. For those attributes you may need to define an inner object or use one of existing in `inner` directory:

* `Fee` – for fee object with fixed and calculated percent amounts
* `XdrEnumValue` – for XDR enum value JSON representation with name and value fields
* `XdrEnumBitmask` – for XDR-enum-based bitmask JSON representation with mask field and flags list

### Resource relation
Resource relation has the following structure:

|Name|Description|
|----|-----------|
|name|JSON name of the relation i.e. `owner` for asset, etc.|
|resource|Key of the resource this is pointing at i.e. `accounts`|
|is_collection|Indicates wether it's a "one-to-one" or "one-to-many" relation|

**Important notice:** Relation must contain the key of the resource it's pointing at in order to generate typed code for consumers. In cases when the relation can point to a set of resource types you have to describe an empty base resource, use it's key for relation and then inherit concrete resources from it.

For example, `Operation` has `details` relation. There are many operation details resources. So create `OperationDetails` resource with `operation-details` key and set it to the relation. Then specify `operation-details` as a `base` for each concrete operation details resource.

## Describing an inner object
Non-resource inner object must be described as a `.yaml` file in `inner` directory.

Inner object specification has the following structure:

|Name|Description|
|----|-----------|
|name|Name of the object i.e. `Fee`, `XdrEnumValue`, etc.|
|attributes|List of object attributes|

Inner object attribute has the same structure as a [resource attribute](#resource-attribute)

## Examples
### Simple resource

```yaml
name: Asset
key: assets
attributes:
- 
  name: pre_issuance_asset_signer
  type: String
-
  name: details
  type: Object
-
  name: issued
  type: Amount
-
  name: policies
  type: XdrEnumBitmask
-
  name: trailing_digits
  type: Int64

relations:
- 
  name: owner
  is_collection: false
  resource: accounts
```
### Inner object

```yaml
name: XdrEnumValue
attributes:
  -
    name: name
    type: String
  -
    name: value
    type: Int32
```

### Resource inheritance

Base resource:

```yaml
name: OperationDetails
key: operation-details
```

Inherited resource (has a `base` field):

```yaml
name: OpCheckSaleStateDetails
key: operations-check-sale-state
base: operation-details
attributes:
  -
    name: effect
    type: XdrEnumValue
relations:
  -
    name: sale
    is_collection: false
    resource: sales
```

Base resource key in relation:

```yaml
name: Operation
key: operations

attributes:
  -
    name: applied_at
    type: Date

relations:
  -
    name: details
    is_collection: false
    resource: operation-details
    
    # Base resource 'operation-details' specified,
    # may point to any concrete operation details
    # resource inherited from it
```

## Validation
Resources and inner objects specifications rules are described in YAML schemas in `schema` directory which can be used for validation. 

To validate all files run `ruby validator.rb` (run `bundle install` for the first time to install Ruby dependencies)

You can also load schema into [jsonschemalint.com](https://jsonschemalint.com/#/version/draft-07/markup/yaml) and use it as an editor with runtime validation.

**Important notice:** resource keys are validated by the set of allowed values. If you are adding a new resource with previously unknown key add this key to the `schema/resource.yaml` schema.