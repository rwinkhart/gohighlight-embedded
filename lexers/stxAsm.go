//go:build stxAsm || stxAll

package lexers

func init() {
	syntaxMap["asm"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: asm
rules:
    - statement: "\\b(?i)(mov|aaa|aad|aam|aas|adc|add|and|call|cbw|clc|cld|cli|cmc|cmp|cmpsb|cmpsw|cwd|daa|das|dec|div|esc|hlt|idiv|imul|in|inc|int|into|iret|ja|jae|jb|jbe|jc|je|jg|jge|jl|jle|jna|jnae|jnb|jnbe|jnc|jne|jng|jnge|jnl|jnle|jno|jnp|jns|jnz|jo|jp|jpe|jpo|js|jz|jcxz|jmp|lahf|lds|lea|les|lock|lodsb|lodsw|loop|loope|loopne|loopnz|loopz|movsb|movsw|mul|neg|nop|or|pop|popf|push|pushf|rcl|rcr|rep|repe|repne|repnz|repz|ret|retn|retf|rol|ror|sahf|sal|sar|sbb|scasb|scasw|shl|shr|stc|std|sti|stosb|stosw|sub|test|wait|xchg|xlat|xor)(?-i)\\b"
    - statement: "\\b(?i)(bound|enter|ins|leave|outs|popa|pusha)(?-i)\\b"
    - statement: "\\b(?i)(arpl|clts|lar|lgdt|lidt|lldt|lmsw|loadall|lsl|ltr|sgdt|sidt|sldt|smsw|str|verr|verw)(?-i)\\b"
    - statement: "\\b(?i)(bsf|bsr|bt|btc|btr|bts|cdq|cmpsd|cwde|insd|iret|iretd|iretf|jecxz|lfs|lgs|lss|lodsd|loopw|loopew|loopnew|loopnzw|loopzw|loopd|looped|loopned|loopnzd|loopzd|cr|tr|dr|movsd|movsx|movzx|outsd|popad|popfd|pushad|pushfd|scasd|seta|setae|setb|setbe|setc|sete|setg|setge|setl|setle|setna|setnae|setnb|setnbe|setnc|setne|setng|setnge|setnl|setnle|setno|setnp|setns|setnz|seto|setp|setpe|setpo|sets|setz|shdl|shrd|stosd)(?-i)\\b"
    - statement: "\\b(?i)(bswap|cmpxcgh|invd|invlpg|wbinvd|xadd)(?-i)\\b"
    - statement: "\\b(?i)(cpuid|cmpxchg8b|rdmsr|rdtsc|wrmsr|rsm)(?-i)\\b"
    - statement: "\\b(?i)(rdpmc)(?-i)\\b"
    - statement: "\\b(?i)(syscall|sysret)(?-i)\\b"
    - statement: "\\b(?i)(cmova|cmovae|cmovb|cmovbe|cmovc|cmove|cmovg|cmovge|cmovl|cmovle|cmovna|cmovnae|cmovnb|cmovnbe|cmovnc|cmovne|cmovng|cmovnge|cmovnle|cmovno|cmovpn|cmovns|cmovnz|cmovo|cmovp|cmovpe|cmovpo|cmovs|cmovz|sysenter|sysexit|ud2)(?-i)\\b"
    - statement: "\\b(?i)(maskmovq|movntps|movntq|prefetch0|prefetch1|prefetch2|prefetchnta|sfence)(?-i)\\b"
    - statement: "\\b(?i)(clflush|lfence|maskmovdqu|mfence|movntdq|movnti|movntpd|pause)(?-i)\\b"
    - statement: "\\b(?i)(monitor|mwait)(?-i)\\b"
    - statement: "\\b(?i)(cdqe|cqo|cmpsq|cmpxchg16b|iretq|jrcxz|lodsq|movsdx|popfq|pushfq|rdtscp|scasq|stosq|swapgs)(?-i)\\b"
    - statement: "\\b(?i)(clgi|invlpga|skinit|stgi|vmload|vmmcall|vmrun|vmsave)(?-i)\\b"
    - statement: "\\b(?i)(vmptrdl|vmptrst|vmclear|vmread|vmwrite|vmcall|vmlaunch|vmresume|vmxoff|vmxon)(?-i)\\b"
    - statement: "\\b(?i)(lzcnt|popcnt)(?-i)\\b"
    - statement: "\\b(?i)(bextr|blcfill|blci|blcic|blcmask|blcs|blsfill|blsic|t1mskc|tzmsk)(?-i)\\b"
    - statement: "\\b(?i)(f2xm1|fabs|fadd|faddp|fbld|fbstp|fchs|fclex|fcom|fcomp|fcompp|fdecstp|fdisi|fdiv|fvidp|fdivr|fdivrp|feni|ffree|fiadd|ficom|ficomp|fidiv|fidivr|fild|fimul|fincstp|finit|fist|fistp|fisub|fisubr|fld|fld1|fldcw|fldenv|fldenvw|fldl2e|fldl2t|fldlg2|fldln2|fldpi|fldz|fmul|fmulp|fnclex|fndisi|fneni|fninit|fnop|fnsave|fnsavenew|fnstcw|fnstenv|fnstenvw|fnstsw|fpatan|fprem|fptan|frndint|frstor|frstorw|fsave|fsavew|fscale|fsqrt|fst|fstcw|fstenv|fstenvw|fstp|fstpsw|fsub|fsubp|fsubr|fsubrp|ftst|fwait|fxam|fxch|fxtract|fyl2x|fyl2xp1)(?-i)\\b"
    - statement: "\\b(?i)(fsetpm)(?-i)\\b"
    - statement: "\\b(?i)(fcos|fldenvd|fsaved|fstenvd|fprem1|frstord|fsin|fsincos|fstenvd|fucom|fucomp|fucompp)(?-i)\\b"
    - statement: "\\b(?i)(fcmovb|fcmovbe|fcmove|fcmove|fcmovnb|fcmovnbe|fcmovne|fcmovnu|fcmovu)(?-i)\\b"
    - statement: "\\b(?i)(fcomi|fcomip|fucomi|fucomip)(?-i)\\b"
    - statement: "\\b(?i)(fxrstor|fxsave)(?-i)\\b"
    - statement: "\\b(?i)(fisttp)(?-i)\\b"
    - statement: "\\b(?i)(ffreep)(?-i)\\b"
    - statement: "\\b(?i)(emms|movd|movq|packssdw|packsswb|packuswb|paddb|paddw|paddd|paddsb|paddsw|paddusb|paddusw|pand|pandn|por|pxor|pcmpeqb|pcmpeqw|pcmpeqd|pcmpgtb|pcmpgtw|pcmpgtd|pmaddwd|pmulhw|pmullw|psllw|pslld|psllq|psrad|psraw|psrlw|psrld|psrlq|psubb|psubw|psubd|psubsb|psubsw|psubusb|punpckhbw|punpckhwd|punpckhdq|punkcklbw|punpckldq|punpcklwd)(?-i)\\b"
    - statement: "\\b(?i)(paveb|paddsiw|pmagw|pdistib|psubsiw|pmwzb|pmulhrw|pmvnzb|pmvlzb|pmvgezb|pmulhriw|pmachriw)(?-i)\\b"
    - statement: "\\b(?i)(femms|pavgusb|pf2id|pfacc|pfadd|pfcmpeq|pfcmpge|pfcmpgt|pfmax|pfmin|pfmul|pfrcp|pfrcpit1|pfrcpit2|pfrsqit1|pfrsqrt|pfsub|pfsubr|pi2fd|pmulhrw|prefetch|prefetchw)(?-i)\\b"
    - statement: "\\b(?i)(pf2iw|pfnacc|pfpnacc|pi2fw|pswapd)(?-i)\\b"
    - statement: "\\b(?i)(pfrsqrtv|pfrcpv)(?-i)\\b"
    - statement: "\\b(?i)(addps|addss|cmpps|cmpss|comiss|cvtpi2ps|cvtps2pi|cvtsi2ss|cvtss2si|cvttps2pi|cvttss2si|divps|divss|ldmxcsr|maxps|maxss|minps|minss|movaps|movhlps|movhps|movlhps|movlps|movmskps|movntps|movss|movups|mulps|mulss|rcpps|rcpss|rsqrtps|rsqrtss|shufps|sqrtps|sqrtss|stmxcsr|subps|subss|ucomiss|unpckhps|unpcklps)(?-i)\\b"
    - statement: "\\b(?i)(andnps|andps|orps|pavgb|pavgw|pextrw|pinsrw|pmaxsw|pmaxub|pminsw|pminub|pmovmskb|pmulhuw|psadbw|pshufw|xorps)(?-i)\\b"
    - statement: "\\b(?i)(movups|movss|movlps|movhlps|movlps|unpcklps|unpckhps|movhps|movlhps|prefetchnta|prefetch0|prefetch1|prefetch2|nop|movaps|cvtpi2ps|cvtsi2ss|cvtps2pi|cvttss2si|cvtps2pi|cvtss2si|ucomiss|comiss|sqrtps|sqrtss|rsqrtps|rsqrtss|rcpps|andps|orps|xorps|addps|addss|mulps|mulss|subps|subss|minps|minss|divps|divss|maxps|maxss|pshufw|ldmxcsr|stmxcsr|sfence|cmpps|cmpss|pinsrw|pextrw|shufps|pmovmskb|pminub|pmaxub|pavgb|pavgw|pmulhuw|movntq|pminsw|pmaxsw|psadbw|maskmovq)(?-i)\\b"
    - statement: "\\b(?i)(addpd|addsd|addnpd|cmppd|cmpsd)(?-i)\\b"
    - statement: "\\b(?i)(addpd|addsd|andnpd|andpd|cmppd|cmpsd|comisd|cvtdq2pd|cvtdq2ps|cvtpd2dq|cvtpd2pi|cvtpd2ps|cvtpi2pd|cvtps2dq|cvtps2pd|cvtsd2si|cvtsd2ss|cvtsi2sd|cvtss2sd|cvttpd2dq|cvttpd2pi|cvttps2dq|cvttsd2si|divpd|divsd|maxpd|maxsd|minpd|minsd|movapd|movhpd|movlpd|movmskpd|movsd|movupd|mulpd|mulsd|orpd|shufpd|sqrtpd|sqrtsd|subpd|subsd|ucomisd|unpckhpd|unpcklpd|xorpd)(?-i)\\b"
    - statement: "\\b(?i)(movdq2q|movdqa|movdqu|movq2dq|paddq|psubq|pmuludq|pshufhw|pshuflw|pshufd|pslldq|psrldq|punpckhqdq|punpcklqdq)(?-i)\\b"
    - statement: "\\b(?i)(addsubpd|addsubps|haddpd|haddps|hsubpd|hsubps|movddup|movshdup|movsldu)(?-i)\\b"
    - statement: "\\b(?i)(lddqu)(?-i)\\b"
    - statement: "\\b(?i)(psignw|psignd|psignb|pshufb|pmulhrsw|pmaddubsw|phsubw|phsubsw|phsubd|phaddw|phaddsw|phaddd|palignr|pabsw|pabsd|pabsb)(?-i)\\b"
    - statement: "\\b(?i)(dpps|dppd|blendps|blendpd|blendvps|blendvpd|roundps|roundss|roundpd|roundsd|insertps|extractps)(?-i)\\b"
    - statement: "\\b(?i)(mpsadbw|phminposuw|pmulld|pmuldq|pblendvb|pblendw|pminsb|pmaxsb|pminuw|pmaxuw|pminud|pmaxud|pminsd|pmaxsd|pinsrb|pinsrd/pinsrq|pextrb|pextrw|pextrd/pextrq|pmovsxbw|pmovzxbw|pmovsxbd|pmovzxbd|pmovsxbq|pmovzxbq|pmovsxwd|pmovzxwd|pmovsxwq|pmovzxwq|pmovsxdq|pmovzxdq|ptest|pcmpeqq|packusdw|movntdqa)(?-i)\\b"
    - statement: "\\b(?i)(extrq|insertq|movntsd|movntss)(?-i)\\b"
    - statement: "\\b(?i)(crc32|pcmpestri|pcmpestrm|pcmpistri|pcmpistrm|pcmpgtq)(?-i)\\b"
    - statement: "\\b(?i)(vfmaddpd|vfmaddps|vfmaddsd|vfmaddss|vfmaddsubpd|vfmaddsubps|vfmsubaddpd|vfmsubaddps|vfmsubpd|vfmsubps|vfmsubsd|vfmsubss|vfnmaddpd|vfnmaddps|vfnmaddsd|vfnmaddss|vfnmsubps|vfnmsubsd|vfnmsubss)(?-i)\\b"
    - statement: "\\b(?i)(aesenc|aesenclast|aesdec|aesdeclast|aeskeygenassist|aesimc)(?-i)\\b"
    - statement: "\\b(?i)(sha1rnds4|sha1nexte|sha1msg1|sha1msg2|sha256rnds2|sha256msg1|sha256msg2)(?-i)\\b"
    - statement: "\\b(?i)(aam|aad|salc|icebp|loadall|loadalld|ud1)(?-i)\\b"
    - identifier: "\\b(?i)(al|ah|bl|bh|cl|ch|dl|dh|bpl|sil|r8b|r9b|r10b|r11b|dil|spl|r12b|r13b|r14b|r15)(?-i)\\b"
    - identifier: "\\b(?i)(cw|sw|tw|fp_ds|fp_opc|fp_ip|fp_dp|fp_cs|cs|ss|ds|es|fs|gs|gdtr|idtr|tr|ldtr|ax|bx|cx|dx|bp|si|r8w|r9w|r10w|r11w|di|sp|r12w|r13w|r14w|r15w|ip)(?-i)\\b"
    - identifier: "\\b(?i)(fp_dp|fp_ip|eax|ebx|ecx|edx|ebp|esi|r8d|r9d|r10d|r11d|edi|esp|r12d|r13d|r14d|r15d|eip|eflags|mxcsr)(?-i)\\b"
    - identifier: "\\b(?i)(mm0|mm1|mm2|mm3|mm4|mm5|mm6|mm7|rax|rbx|rcx|rdx|rbp|rsi|r8|r9|r10|r11|rdi|rsp|r12|r13|r14|r15|rip|rflags|cr0|cr1|cr2|cr3|cr4|cr5|cr6|cr7|cr8|cr9|cr10|cr11|cr12|cr13|cr14|cr15|msw|dr0|dr1|dr2|dr3|r4|dr5|dr6|dr7|dr8|dr9|dr10|dr11|dr12|dr13|dr14|dr15)(?-i)\\b"
    - identifier: "\\b(?i)(st0|st1|st2|st3|st4|st5|st6|st7)(?-i)\\b"
    - identifier: "\\b(?i)(xmm0|xmm1|xmm2|xmm3|xmm4|xmm5|xmm6|xmm7|xmm8|xmm9|xmm10|xmm11|xmm12|xmm13|xmm14|xmm15)(?-i)\\b"
    - identifier: "\\b(?i)(ymm0|ymm1|ymm2|ymm3|ymm4|ymm5|ymm6|ymm7|ymm8|ymm9|ymm10|ymm11|ymm12|ymm13|ymm14|ymm15)(?-i)\\b"
    - identifier: "\\b(?i)(zmm0|zmm1|zmm2|zmm3|zmm4|zmm5|zmm6|zmm7|zmm8|zmm9|zmm10|zmm11|zmm12|zmm13|zmm14|zmm15|zmm16|zmm17|zmm18|zmm19|zmm20|zmm21|zmm22|zmm23|zmm24|zmm25|zmm26|zmm27|zmm28|zmm29|zmm30|zmm31)(?-i)\\b"
    - constant.number: "\\b(|h|A|0x)+[0-9]+(|h|A)+\\b"
    - constant.number: "\\b0x[0-9 a-f A-F]+\\b"
    - preproc: "%+(\\+|\\?|\\?\\?|)[a-z A-Z 0-9]+"
    - preproc: "%\\[[. a-z A-Z 0-9]*\\]"
    - statement: "\\b(?i)(extern|global|section|segment|_start|\\.text|\\.data|\\.bss)(?-i)\\b"
    - statement: "\\b(?i)(db|dw|dd|dq|dt|ddq|do)(?-i)\\b"
    - identifier: "[a-z A-Z 0-9 _]+:"
    - constant.string:
        start: "\""
        end: "\""
        skip: "\\\\."
        rules:
            - constant.specialChar: "\\\\."
    - constant.string:
        start: "'"
        end: "'"
        skip: "\\\\."
        rules:
            - constant.specialChar: "\\\\."
    - comment:
        start: ";"
        end: "$"
        rules:
            - todo: "(TODO|XXX|FIXME):?"`)
	}}
}
