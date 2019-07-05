# dotenv

[![GoDoc](https://godoc.org/github.com/radlinskii/dotenv?status.svg)](https://godoc.org/github.com/radlinskii/dotenv)
[![Build Status](https://travis-ci.com/radlinskii/dotenv.svg?branch=master)](https://travis-ci.com/radlinskii/dotenv)
[![Go Report Card](https://goreportcard.com/badge/github.com/radlinskii/dotenv)](https://goreportcard.com/report/github.com/radlinskii/dotenv)
[![version](https://img.shields.io/github/release/radlinskii/dotenv.svg)](https://img.shields.io/github/release/radlinskii/dotenv.svg)

Tiny library for setting environment variables specified in `.env` files.

## Supported .env Syntax

1. BASIC_KEY=basic_value
2. WHITE_SPACES = are trimmed
3. \# lines starting with "#" are omitted
4. \# blank lines are omitted
5. ALREADY_EXPORTED_VARIABLES="are not overwritten"
