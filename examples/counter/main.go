package main

import (
	"fmt"
	"strings"

	"github.com/marcsantiago/collections"
	"github.com/marcsantiago/collections/counter"
)

var text = `Lorem ipsum dolor sit amet, nec insolens intellegebat ea. Usu inani simul et, est laudem fabulas consulatu eu. Putent principes ei sed. Volumus reprimique consectetuer an sed, nisl sale recteque ne mei. Qualisque consulatu te nec.
Eu erat omittam mel, at dicant electram theophrastus vix. Ei sit diceret nusquam albucius. Qui rebum exerci tacimates eu. Sea ludus vitae noluisse in, an vis impetus sanctus impedit. Vix ad explicari vulputate.
In est aperiam docendi constituto, mea adversarium conclusionemque ne, ad usu prima dictas epicuri. Ex nec hinc etiam antiopam, nostro iuvaret te nam, cu mea augue homero. Nam sanctus maiestatis ne. Esse sanctus ea mel, senserit disputando ut duo, quot persius interesset has ut. Qui nihil veniam te, ad soleat debitis fabellas eam, no vix sale probo.
Per cu scaevola temporibus, eos mazim everti electram ad, id sale exerci mei. Sit eu utamur fastidii explicari, mel ei alii nullam dissentiet. Iusto tacimates hendrerit nam et, esse omnes sententiae at ius. Has ei errem sensibus vulputate, at aliquid adversarium sit. Est case tantas signiferumque ei, mel ex epicurei patrioque.
Assum patrioque philosophia eu vim, nobis sensibus id sea. Usu id virtute luptatum, everti delicata te sit. Mei id regione salutandi, ut habeo dissentiet his, illud consulatu ad eum. Mel nonumy regione ei, eum omnium efficiendi ne, denique eligendi ut sed. Dico suas delicata est ea, mei enim iudico suscipiantur ne. Nam ullum delicata ut, ius ei tale rebum. Sed persius molestie electram an.`

func main() {
	words := collections.StringValues(strings.Split(text, " "))
	count := counter.Counter(words.Data())
	// optionally I could pass in the underline data type to avoid reflection
	// count := counter.Counter(words.Data(), collections.StringSliceType)
	fmt.Println("Most All", count.MostCommon(-1))
	fmt.Println("Top ten", count.MostCommon(10))
}
