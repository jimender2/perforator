#pragma once

#ifdef __GNUC__
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wunused-parameter"
#endif

//===-- MCTargetOptionsCommandFlags.h --------------------------*- C++ -*-===//
//
// Part of the LLVM Project, under the Apache License v2.0 with LLVM Exceptions.
// See https://llvm.org/LICENSE.txt for license information.
// SPDX-License-Identifier: Apache-2.0 WITH LLVM-exception
//
//===----------------------------------------------------------------------===//
//
// This file contains machine code-specific flags that are shared between
// different command line tools.
//
//===----------------------------------------------------------------------===//

#ifndef LLVM_MC_MCTARGETOPTIONSCOMMANDFLAGS_H
#define LLVM_MC_MCTARGETOPTIONSCOMMANDFLAGS_H

#include <optional>
#include <string>

namespace llvm {

class MCTargetOptions;
enum class EmitDwarfUnwindType;

namespace mc {

bool getRelaxAll();
std::optional<bool> getExplicitRelaxAll();

bool getIncrementalLinkerCompatible();

int getDwarfVersion();

bool getDwarf64();

EmitDwarfUnwindType getEmitDwarfUnwind();

bool getEmitCompactUnwindNonCanonical();

bool getShowMCInst();

bool getFatalWarnings();

bool getNoWarn();

bool getNoDeprecatedWarn();

bool getNoTypeCheck();

std::string getABIName();

std::string getAsSecureLogFile();

/// Create this object with static storage to register mc-related command
/// line options.
struct RegisterMCTargetOptionsFlags {
  RegisterMCTargetOptionsFlags();
};

MCTargetOptions InitMCTargetOptionsFromFlags();

} // namespace mc

} // namespace llvm

#endif

#ifdef __GNUC__
#pragma GCC diagnostic pop
#endif
