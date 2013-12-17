# Seisan

Seisan solution for small team.

## Installation

You need a few steps to setup seisan data repository. `example` directory in `enishitech/seisan` will be a good reference for you.

Put `Gemfile`:

```
source 'https://rubygems.org'

gem 'seisan'
```

Run bundler:

```shell
% bundle
```

Put `Rakefile`:

```
require 'seisan/task'
```
Create `data` directory to store data.

```shell
% mkdir data
```

OK, now everything is set up. Run

```shell
% bundle exec rake seisan target=2013/07
```

Then you will have an empty monthly report (because you have no record in seisan data) at `output/2013-07.xlsx`.

You may want to add `output` directory to `.gitignore`:

```shell
% cat .gitignore
/output
```

## Usage

Given you record seisan data in the following format:

```
applicant: 佐藤
expense:
  - date: 2013-6-24
    amount: 105
    remarks: 電池代

  - date: 2013-7-4
    amount: 2080
    remarks: JR代
```

And you file seisan data as git repository like this:

```
% tree
data
└── 2013
    ├── ...
    ├── 07
    │   ├── 20130709-shidara.yaml
    │   ├── 20130711-sato.yaml
    │   ├── 20130712-sato.yaml
    │   ├── 20130717-shidara.yaml
    │   └── 20130719-shimada.yaml
    └── 08
        ├── 20130802-sato.yaml
        ├── 20130802-sekiya.yaml
        ├── 20130803-shidara.yaml
        └── 08-shidara.yaml
```

Put `Rakefile` to your seisan data repository,

Then you can generate seisan report.

```shell
% bundle exec rake seisan target=2013/07
```

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

-----

&copy; 2013 Enishi Tech Inc.
