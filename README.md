# dotenv

[![GoDoc](https://godoc.org/github.com/radlinskii/dotenv?status.svg)](https://godoc.org/github.com/radlinskii/dotenv)
[![Build Status](https://travis-ci.com/radlinskii/dotenv.svg?branch=master)](https://travis-ci.com/radlinskii/dotenv)
[![Go Report Card](https://goreportcard.com/badge/github.com/radlinskii/dotenv)](https://goreportcard.com/report/github.com/radlinskii/dotenv)
[![version](https://img.shields.io/github/release/radlinskii/dotenv.svg)](https://img.shields.io/github/release/radlinskii/dotenv.svg)

Tiny library for setting environment variables specified in `.env` files.

## Supported .env Syntax

1. Variables should be stored as Key-Value pairs separated by "=" sign.
2. Whitespaces around key and value are trimmed.
3. Lines starting with "#" sign are omitted.
4. Blank lines are omitted.
5. Environment variables that are already exported are not getting overwritten by those read from .env file.
