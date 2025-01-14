import {
  Box,
  Stack,
  Table,
  TableContainer,
  Tbody,
  Td,
  Tr,
  VStack,
} from "@chakra-ui/react";
import {
  Accordion,
  AccordionButton,
  AccordionIcon,
  AccordionItem,
  AccordionPanel,
  Skeleton,
  Text,
} from "@chakra-ui/react";
import React, { useMemo } from "react";
import {
  useGetUser,
  useListRequestEvents,
} from "../utils/backend-client/end-user/end-user";
import { RequestDetail } from "../utils/backend-client/types";
import { renderTiming } from "../utils/renderTiming";
import { CFTimelineRow } from "./CFTimelineRow";
export const AuditLog: React.FC<{ request?: RequestDetail }> = ({
  request,
}) => {
  const { data } = useListRequestEvents(request?.id || "", {
    swr: {
      refreshInterval: 5000,
    },
  });

  const events = useMemo(() => {
    const items: JSX.Element[] = [];
    // use map here to ensure order is preserved
    // foreach is not synchronous
    const l = data?.events.length || 0;
    data?.events.forEach((e, i) => {
      if (e.grantCreated) {
        items.push(
          <CFTimelineRow
            arrLength={l}
            header={<Text>Grant created</Text>}
            index={i}
            key={i}
            timestamp={new Date(e.createdAt)}
          />
        );
      } else if (e.fromGrantStatus && e.actor) {
        let selectCase: string | undefined;
        if (e.toGrantStatus === "ACTIVE") selectCase = " approved the request";
        if (e.toGrantStatus === "REVOKED") selectCase = " revoked the grant";
        items.push(
          <CFTimelineRow
            arrLength={l}
            header={
              <Text>
                <UserText userId={e.actor || ""} />{" "}
                {
                  /** if null, fallback to default case */
                  selectCase ??
                    `changed grant status from
              ${e.fromGrantStatus?.toLowerCase()} to ${e.toGrantStatus?.toLowerCase()}`
                }
              </Text>
            }
            index={i}
            key={i}
            timestamp={new Date(e.createdAt)}
          />
        );
      } else if (e.fromGrantStatus && e.grantFailureReason) {
        items.push(
          <CFTimelineRow
            arrLength={l}
            header={
              <>
                <Text>Grant failed due to an error:</Text>
                <Accordion allowToggle>
                  <AccordionItem borderStyle="none">
                    <h2>
                      <AccordionButton>
                        <Text
                          fontSize="sm"
                          color="#757575"
                          fontWeight="normal"
                          flex="1"
                          textAlign="left"
                        >
                          Details
                        </Text>
                        <AccordionIcon />
                      </AccordionButton>
                    </h2>
                    <AccordionPanel pb={4}>
                      <Text
                        fontSize="sm"
                        fontWeight="normal"
                        flex="1"
                        textAlign="left"
                      >
                        {e.grantFailureReason}
                      </Text>
                    </AccordionPanel>
                  </AccordionItem>
                </Accordion>
              </>
            }
            index={i}
            key={i}
            timestamp={new Date(e.createdAt)}
          />
        );
      } else if (e.fromGrantStatus) {
        items.push(
          <CFTimelineRow
            arrLength={l}
            header={
              <Text>
                {`Grant status changed from ${e.fromGrantStatus?.toLowerCase()} to
              ${e.toGrantStatus?.toLowerCase()}`}
              </Text>
            }
            index={i}
            key={i}
            timestamp={new Date(e.createdAt)}
          />
        );
      } else if (e.fromTiming && e.actor) {
        items.push(
          <CFTimelineRow
            arrLength={l}
            header={
              <Text>
                <UserText userId={e.actor || ""} />
                {` changed request timing from
              ${renderTiming(e.fromTiming)} to ${renderTiming(e.toTiming)}`}
              </Text>
            }
            index={i}
            key={i}
            timestamp={new Date(e.createdAt)}
          />
        );
      } else if (e.fromStatus && e.actor) {
        items.push(
          <CFTimelineRow
            arrLength={l}
            header={
              <Text>
                <UserText userId={e.actor || ""} />
                {` ${e.toStatus?.toLowerCase()} the request`}
              </Text>
            }
            index={i}
            key={i}
            timestamp={new Date(e.createdAt)}
          />
        );
      } else if (e.fromStatus?.toLowerCase()) {
        items.push(
          <CFTimelineRow
            arrLength={l}
            header={
              <Text>
                {`Granted Approvals changed request status from
              ${e.fromStatus?.toLowerCase()} to ${e.toStatus?.toLowerCase()}`}
              </Text>
            }
            index={i}
            key={i}
            timestamp={new Date(e.createdAt)}
          />
        );
      } else if (e.requestCreated) {
        items.push(
          <CFTimelineRow
            arrLength={l}
            header={
              <Text>
                {`Request created by `}
                <UserText userId={e.actor || ""} />
              </Text>
            }
            index={i}
            key={i}
            timestamp={new Date(e.createdAt)}
          />
        );
      } else if (e.recordedEvent) {
        items.push(
          <CFTimelineRow
            arrLength={l}
            header={
              <Stack w="100%">
                <Text>
                  {`Action performed by `}
                  <UserText userId={e.actor || ""} />
                </Text>
                <TableContainer>
                  <Table size="sm" variant={"unstyled"}>
                    <Tbody>
                      {Object.entries(e.recordedEvent).map(([k, v]) => (
                        <Tr key={k} color="GrayText">
                          <Td
                            w="20%"
                            borderColor="gray.200"
                            borderWidth={"1px"}
                            fontWeight="thin"
                          >
                            {k}
                          </Td>
                          <Td borderColor="gray.200" borderWidth={"1px"}>
                            {v}
                          </Td>
                        </Tr>
                      ))}
                    </Tbody>
                  </Table>
                </TableContainer>
              </Stack>
            }
            index={i}
            timestamp={new Date(e.createdAt)}
          />
        );
      }
    });
    return items;
  }, [data]);
  if (!request || data === undefined) {
    return (
      <VStack flex={1} align="left">
        <Box textStyle="Heading/H4" as="h4" mb={8}>
          Audit Log
        </Box>
        <Skeleton h={30} w="100%" />
      </VStack>
    );
  }

  return (
    <VStack flex={1} align="left">
      <Box textStyle="Heading/H4" as="h4" mb={8}>
        Audit Log
      </Box>
      {events}
    </VStack>
  );
};

const UserText: React.FC<{ userId: string }> = ({ userId }) => {
  const { data } = useGetUser(userId);
  if (!data) {
    return <Text></Text>;
  }
  if (data.firstName && data.lastName) {
    <Text>{`${data.firstName} ${data.lastName}`}</Text>;
  }
  return <Text>{data.email}</Text>;
};
