
/**

 * @kind path-problem
 * @id go/lfi
 * @severity error

 */

import go

module MyFlowConfiguration implements DataFlow::ConfigSig {
	predicate isSource(DataFlow::Node source) {
		exists(Function qp, CallExpr call |
			qp.hasQualifiedName("github.com/labstack/echo.Context", "QueryParam") and
			call.getTarget() = qp and
			source = call.getTarget().getACall().getResult()
		) 
	}

	predicate isSink(DataFlow::Node sink) {
		exists(Function osOpen, CallExpr call |
			osOpen.hasQualifiedName("os", "Open") and
			call.getTarget() = osOpen and
			sink.asExpr() = call.getArgument(0)
		)
	}
}

module MyFlow = TaintTracking::Global<MyFlowConfiguration>;
import MyFlow::PathGraph

from DataFlow::Node source, DataFlow::Node sink
where MyFlow::flow(source, sink)
select sink.asExpr(), source, sink, "Potential Local File Inclusion (LFI) vulnerability detected $@.", sink, "here"
