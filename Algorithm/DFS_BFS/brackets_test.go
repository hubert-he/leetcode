package DFS_BFS

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBraceExpansionII(t *testing.T){
	for caseId, testCase := range []struct{
		s		string
		want	[]string
	}{
		{"{a,b}{c,{d,e}}", []string{"ac","ad","ae","bc","bd","be"}},
		{"{a,b}c{d,e}f", []string{"acdf","acef","bcdf","bcef"}},
		{"{{a,z},a{b,c},{ab,z}}", []string{"a","ab","ac","z"}},
	}{
		result := BraceExpansionII_2(testCase.s)
		if !assert.Equal(t, testCase.want, result, "case-%d Failed: result=%v but want=%v", caseId, result, testCase.want){
			break
		}
	}
}

func TestRemoveInvalidParentheses(t *testing.T){
	for caseId, testCase := range []struct{
		s		string
		want	[]string
	}{
		{"()", []string{"()"}},
		{")(", []string{""}},
		{")(f", []string{"f"}}, // 检验去重
		{"(a)())()", []string{"(a())()","(a)()()"}},
		{"()())()", []string{"(())()","()()()"}},
		{"((()((s((((()", []string{"()s()","()(s)","(()s)"}}, // 超时case
	}{
		result := removeInvalidParentheses_BFS(testCase.s)
		if !assert.ElementsMatch(t, testCase.want, result, "case-%d Failed: result=%v but want=%v", caseId, result, testCase.want){
			break
		}
	}
}

func TestReverseParentheses(t *testing.T) {
	for caseId, testCase := range []struct{
		s		string
		want	string
	}{
		{"ta()usw((((a))))", "tauswa"},
		{"(abcd)", "dcba"},
		{"(u(love)i)", "iloveu"},
		{"(ed(et(oc))el)", "leetcode"},
		{"(a)(b)", "ab"},
		{"((e(f()(((()vbl))s)i(hbo)(j((emr(g((dyvz(j(k))qn(r(s))(by()lg)(z)(v))po(ri))uq)(())(i)((((joovdi(r(hov)tk)ycpv))(uu)n)((pc(kmhzko(h)j())x)idpe(tf(a()j)lcszq)el)e)(q)s)))h((()hq))ty)z((r)(etuimhqk(vc)o(x(eavtr)c())gr(iaeh))(uijw)ribmj((nmxndbljlphzisqms)q)hp(()(((k)y(qfjwg)t(v)rye(mm)jonu()gwv))(()dtc(nf)a)q)(l(g)ls)(elxperab)ugnutxcd)ucbet)joc(e(ka)))ayudqadlo()(s(rkyp)u)uoukgnkbxvgqpm()u()ofcoobafiyurfx()bwcnlgnjieh((up)lfo(nfzid(wpcttauya((d)(lt(s)fa(o))it(gn)(imb(rp(b)v(w((kt())qcia(lsu(nx)biucqc)g(rjvzm)))(af(p)(km)c(ozd)i(a(ufpmqyty)gd(unoo()ncwc)b(buj)s))(z))yh)goq(u)((kn)kpa)kfe(r(aurgx)ke)xpa(lofufr((r())d)(wlw(ew)))))(k)()(lwq)wksx))pavt(w)n(jn)gewybef(t)djbyk((b))lkqiyxo(on)yckdkzfmradisc()(o)qdl((asms(c((t)zwcc)g)(pc)()e((((rlm(jegb)zcu((bw(mbps(g)n)fgkjkb(fp(vm(tzsp)(t((t)d()c(x(ktviam(e((r)xfktm)vc(w)hi(ylenyelvde(lu((xce)ofbiv)je)t((oqp)jng)vr)o)ctdkzogm((km)nk))gv)xjueo(qmclm()r)ttg))c)wv())qlrs)(sl(fo(e))nxsjmgxt)nyg(dn()((((ri(b(as()qyg)amcy)vk()()((ny(()(x(gu(q)lxx(m()))))))))))))))))))",
			"eakcojeytqhhjemrquvzbylgrsnqjkzvydpoirgiexkmhzkohjcpidpeqzsclajfteljoovdikthovrycpvuunqshboivblsfzdcxtunguelxperablglskygwjfqtvryemmjonugwvanfctdqphsmqsizhpljlbdnxmnqjmbiruijwetuimhqkcvoceavtrxgrheairucbetayudqadlourkypsuoukgnkbxvgqpmuofcoobafiyurfxbwcnlgnjiehnfziddoafstlitnghyrpbvtkqciacqcuibnxuslgmzvjrwaytyqmpfugdcwcnoonubjubsiozdckmpfazbmigoquapkknkfeekaurgxrxpawlwwerdrfufolayuattcpwkqwlwksxofluppavtwnnjgewybeftdjbykblkqiyxonoyckdkzfmradiscoqdlasmsgtzwccccperlmbgejzcudnriycmaasqygbvknyxmxxlquggynsleofnxsjmgxtbwngspbmfgkjkbvwvmpszttdcvgktviamoylenyelvdeejecxofbivultgnjoqpvrihwcvrxfktmectdkzogmknkmxxjueormlcmqttgtcpfqlrs"},
	}{
		result := ReverseParentheses_DFS(testCase.s)
		if !assert.Equal(t, testCase.want, result, "case-%d Failed: result=%v but want=%v",
			caseId, result, testCase.want){
			break
		}
	}

}