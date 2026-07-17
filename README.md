# go-brazil

[![build](https://github.com/flavioltonon/go-brazil/actions/workflows/build.yml/badge.svg)](https://github.com/flavioltonon/go-brazil/actions/workflows/build.yml)
[![Coverage Status](https://coveralls.io/repos/github/flavioltonon/go-brazil/badge.svg?branch=main)](https://coveralls.io/github/flavioltonon/go-brazil?branch=main)

Go Brazil is a library for evaluation of brazilian documents and other patterns.

## Install

```
go get github.com/flavioltonon/go-brazil
```

## Usage

In general, the lib usage involves the generation and evaluation of brazilian document numbers and other brazilian patterns, such as mobile numbers.
Here are the documents/patterns which have been contemplated so far:

- **CNPJ (Cadastro Nacional de Pessoas Jurídicas)** - including the alphanumeric format introduced by IN RFB n° 2.229/2024, via `ParseAlphanumericCNPJ` (positions 0-11 accept `0-9`/`A-Z`, check digits remain numeric; legacy numeric CNPJs are also accepted)
- **CPF (Cadastro de Pessoas Físicas)**
- **Título de Eleitor**
- **PIS**
- **SUS**
- **Mobile numbers**
- **Dates**

 Some usage examples can be found at example/example.go.
