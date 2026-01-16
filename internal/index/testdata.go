package index

import "perimeter/internal/types"

var scanSourceFileTests = []struct {
	Name     string
	Code     string
	Expected []types.SignatureHit
}{
	{
		Name: "express route get",
		Code: "app.get('/test', (req, res) => { res.send('Hello, world!'); })",
		Expected: []types.SignatureHit{
			{
				Path:          "", // Will be set in test
				LineNumber:    1,
				SignatureType: types.ExpressRoute,
			},
		},
	},
	{
		Name: "express route get multiline",
		Code: `app.get('/api/users', (req, res) => {
  const userId = req.params.id;
  res.json({ id: userId, name: 'John' });
})`,
		Expected: []types.SignatureHit{
			{
				Path:          "",
				LineNumber:    1,
				SignatureType: types.ExpressRoute,
			},
		},
	},
	{
		Name: "express route get second line",
		Code: `console.log('Hello, world!');
app.get('/api/users', (req, res) => {
  const userId = req.params.id;
  res.json({ id: userId, name: 'John' });
})`,
		Expected: []types.SignatureHit{
			{
				Path:          "",
				LineNumber:    2,
				SignatureType: types.ExpressRoute,
			},
		},
	},
	{
		Name:     "no express route - plain function",
		Code:     "function getData() { return data; }",
		Expected: []types.SignatureHit{},
	},
	{
		Name:     "no express route - variable assignment",
		Code:     "const get = require('lodash').get;",
		Expected: []types.SignatureHit{},
	},
	{
		Name:     "no express route - empty file",
		Code:     "",
		Expected: []types.SignatureHit{},
	},
	{
		Name:     "no express route - regular code",
		Code:     "const data = { Name: 'test' };",
		Expected: []types.SignatureHit{},
	},
	{
		Name:     "no express route - import statement",
		Code:     "import { get } from 'lodash';",
		Expected: []types.SignatureHit{},
	},
}

var expandSignatureHitSpanTests = []struct {
	Name     string
	Hit      types.SignatureHit
	Code     string
	Expected types.SignatureSpan
}{
	{
		Name: "express route get",
		Hit: types.SignatureHit{
			Path:          "test.js",
			LineNumber:    1,
			SignatureType: types.ExpressRoute,
		},
		Code: "app.get('/test', (req, res) => { res.send('Hello, world!'); })",
		Expected: types.SignatureSpan{
			Path:      "",
			StartLine: 1,
			EndLine:   1,
			Content:   "app.get('/test', (req, res) => { res.send('Hello, world!'); })",
		},
	},
	{
		Name: "express route get multiline",
		Hit: types.SignatureHit{
			Path:          "test.js",
			LineNumber:    1,
			SignatureType: types.ExpressRoute,
		},
		Code: `app.get('/api/users', (req, res) => {
  const userId = req.params.id;
  res.json({ id: userId, name: 'John' });
})`,
		Expected: types.SignatureSpan{
			Path:      "test.js",
			StartLine: 1,
			EndLine:   4,
			Content: `app.get('/api/users', (req, res) => {
  const userId = req.params.id;
  res.json({ id: userId, name: 'John' });
})`,
		},
	},
	{
		Name: "express route get second line",
		Hit: types.SignatureHit{
			Path:          "test.js",
			LineNumber:    2,
			SignatureType: types.ExpressRoute,
		},
		Code: `console.log('Hello, world!');
app.get('/api/users', (req, res) => {
  const userId = req.params.id;
  res.json({ id: userId, name: 'John' });
})`,
		Expected: types.SignatureSpan{
			Path:      "test.js",
			StartLine: 2,
			EndLine:   5,
			Content: `app.get('/api/users', (req, res) => {
  const userId = req.params.id;
  res.json({ id: userId, name: 'John' });
})`,
		},
	},
}
