module github.com/codedbypm/jaspergify

go 1.13

replace github.com/codedbypm/jaspergify/entry => /Users/paolodev/Projects/jaspergify/entry

replace github.com/codedbypm/jaspergify/fetch => /Users/paolodev/Projects/jaspergify/fetch

replace github.com/codedbypm/jaspergify/decode => /Users/paolodev/Projects/jaspergify/decode

replace github.com/codedbypm/jaspergify/model => /Users/paolodev/Projects/jaspergify/model

replace github.com/codedbypm/jaspergify/log => /Users/paolodev/Projects/jaspergify/log

require (
	github.com/GoogleCloudPlatform/functions-framework-go v1.0.1
	github.com/codedbypm/jaspergify/decode v0.0.0-20200717212615-7190fa11aa88
	github.com/codedbypm/jaspergify/entry v0.0.0-20200717212615-7190fa11aa88
	github.com/codedbypm/jaspergify/fetch v0.0.0-20200620140609-e90fd1b522b2
	github.com/codedbypm/jaspergify/log v0.0.0-00010101000000-000000000000
	github.com/codedbypm/jaspergify/model v0.0.0-00010101000000-000000000000 // indirect
)
