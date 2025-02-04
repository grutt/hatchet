import { useState, useEffect } from 'react';
import ReactFlow, {
  useNodesState,
  useEdgesState,
  Position,
  MarkerType,
  Node,
  Edge,
} from 'reactflow';
import 'reactflow/dist/style.css';
import StepRunNode from './step-run-node';
import { StepRun, StepRunStatus, WorkflowRun } from '@/lib/api';
import dagre from 'dagre';
import { StepRunPlayground } from './step-run-playground';
import invariant from 'tiny-invariant';

const initBgColor = '#050c1c';

const connectionLineStyle = { stroke: '#fff' };
const nodeTypes = {
  stepNode: StepRunNode,
};

const WorkflowRunVisualizer = ({
  workflowRun,
}: {
  workflowRun: WorkflowRun;
}) => {
  const [nodes, setNodes, onNodesChange] = useNodesState([]);
  const [edges, setEdges, onEdgesChange] = useEdgesState([]);
  const [selectedStepRun, setSelectedStepRun] = useState<StepRun | null>(null);
  const [bgColor] = useState(initBgColor);
  const dagreGraph = new dagre.graphlib.Graph();
  dagreGraph.setDefaultEdgeLabel(() => ({}));

  useEffect(() => {
    const stepEdges =
      workflowRun.jobRuns
        ?.map((job) => {
          invariant(job.stepRuns, 'has stepRuns');
          return job.stepRuns
            .map((stepRun: StepRun) => {
              invariant(stepRun.step, 'has step');

              return (
                stepRun.step.parents
                  ?.map((parent) => {
                    invariant(stepRun.step, 'has step');

                    return {
                      id: `${parent}-${stepRun.step.metadata.id}`,
                      source: parent,
                      target: stepRun.step.metadata.id,
                      animated: stepRun.status === StepRunStatus.RUNNING,
                      style: { stroke: '#fff' },
                      markerEnd: {
                        type: MarkerType.ArrowClosed,
                      },
                    };
                  })
                  .flat() || []
              );
            })
            .flat();
        })
        .flat() || [];

    const stepNodes =
      workflowRun.jobRuns
        ?.map((jobRun) => {
          invariant(jobRun.stepRuns, 'has stepRuns');

          return jobRun.stepRuns.map((stepRun) => {
            invariant(stepRun.step, 'has step');
            const hasChild = stepEdges.some((edge) => {
              invariant(stepRun.step, 'has step');
              return edge?.source === stepRun.step.metadata.id;
            });
            const hasParent =
              stepRun.step?.parents?.length && stepRun.step.parents.length > 0;

            return {
              id: stepRun.step.metadata.id,
              selectable: false,
              type: 'stepNode',
              position: { x: 0, y: 0 }, // positioning gets set by dagre later
              data: {
                stepRun: stepRun,
                onClick: () => {
                  setSelectedStepRun(stepRun);
                },
                variant:
                  hasParent && hasChild
                    ? 'default'
                    : hasChild
                      ? 'output_only'
                      : 'input_only',
              },
            };
          });
        })
        .flat() || [];

    setNodes(stepNodes);
    setEdges(stepEdges);
  }, [workflowRun, setNodes, setEdges]);

  const nodeWidth = 230;
  const nodeHeight = 70;

  const getLayoutedElements = (
    nodes: Node[],
    edges: Edge[],
    direction = 'LR',
  ) => {
    const isHorizontal = direction === 'LR';
    dagreGraph.setGraph({ rankdir: direction });

    nodes.forEach((node) => {
      dagreGraph.setNode(node.id, { width: nodeWidth, height: nodeHeight });
    });

    edges.forEach((edge) => {
      dagreGraph.setEdge(edge.source, edge.target);
    });

    dagre.layout(dagreGraph);

    nodes.forEach((node) => {
      const nodeWithPosition = dagreGraph.node(node.id);
      node.targetPosition = isHorizontal ? Position.Left : Position.Top;
      node.sourcePosition = isHorizontal ? Position.Right : Position.Bottom;

      // We are shifting the dagre node position (anchor=center center) to the top left
      // so it matches the React Flow node anchor point (top left).
      node.position = {
        x: nodeWithPosition.x - nodeWidth / 2,
        y: nodeWithPosition.y - nodeHeight / 2,
      };

      return node;
    });

    return { nodes, edges };
  };

  const dagrLayout = getLayoutedElements(nodes, edges);

  const dagrNodes = dagrLayout.nodes;
  const dagrEdges = dagrLayout.edges;

  return (
    <>
      <ReactFlow
        nodes={dagrNodes}
        edges={dagrEdges}
        onNodesChange={onNodesChange}
        onEdgesChange={onEdgesChange}
        style={{ background: bgColor }}
        nodeTypes={nodeTypes}
        connectionLineStyle={connectionLineStyle}
        snapToGrid={true}
        fitView
        proOptions={{
          hideAttribution: true,
        }}
        className="border-1 border-gray-800 rounded-lg"
        maxZoom={1}
      />
      <StepRunPlayground
        stepRun={selectedStepRun}
        setStepRun={setSelectedStepRun}
      />
    </>
  );
};

export default WorkflowRunVisualizer;
