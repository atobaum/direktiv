import { Dialog, DialogContent, DialogTrigger } from "../../../design/Dialog";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "../../../design/Dropdown";
import { FC, useEffect, useState } from "react";
import {
  Folder,
  FolderUp,
  MoreVertical,
  Play,
  TextCursorInput,
  Trash,
} from "lucide-react";

import Button from "../../../design/Button";
import Delete from "./Delete";
import ExplorerHeader from "./Header";
import { Link } from "react-router-dom";
import { NodeSchemaType } from "../../../api/tree/schema";
import Rename from "./Rename";
import { analyzePath } from "../../../util/router/utils";
import moment from "moment";
import { pages } from "../../../util/router/pages";
import { useListDirectory } from "../../../api/tree/query/get";
import { useNamespace } from "../../../util/store/namespace";

const ExplorerPage: FC = () => {
  const namespace = useNamespace();
  const { path } = pages.explorer.useParams();
  const { data } = useListDirectory({ path });
  const { parent, isRoot } = analyzePath(path);
  const [dialogOpen, setDialogOpen] = useState(false);

  // we only want to use one dialog component for the whole list,
  // so when the user clicks on the delete button in the list, we
  // set the pointer to that node for the dialog
  const [deleteNode, setDeleteNode] = useState<NodeSchemaType>();
  const [renameNode, setRenameNode] = useState<NodeSchemaType>();

  useEffect(() => {
    if (dialogOpen === false) {
      setDeleteNode(undefined);
      setRenameNode(undefined);
    }
  }, [dialogOpen]);

  if (!namespace) return null;

  return (
    <div>
      <ExplorerHeader />
      <div className="flex flex-col space-y-5 p-5 text-sm">
        <div className="flex flex-col space-y-5 ">
          <Dialog open={dialogOpen} onOpenChange={setDialogOpen}>
            {!isRoot && (
              <Link
                to={pages.explorer.createHref({
                  namespace,
                  path: parent?.absolute,
                })}
                className="flex items-center space-x-3"
              >
                <FolderUp className="h-5" />
                <span>..</span>
              </Link>
            )}
            {data?.children?.results.map((file) => {
              const Icon = file.expandedType === "workflow" ? Play : Folder;

              const linkTarget =
                file.expandedType === "workflow"
                  ? pages.workflow.createHref({
                      namespace,
                      path: file.path,
                    })
                  : pages.explorer.createHref({
                      namespace,
                      path: file.path,
                    });

              return (
                <div key={file.name}>
                  <div className="flex items-center space-x-3">
                    <Icon className="h-5" />
                    <Link to={linkTarget} className="flex flex-1">
                      <span className="flex-1">{file.name}</span>
                      <span className="text-gray-8 dark:text-gray-dark-8">
                        {moment(file.updatedAt).fromNow()}
                      </span>
                    </Link>
                    <DropdownMenu>
                      <DropdownMenuTrigger asChild>
                        <Button
                          variant="ghost"
                          size="sm"
                          onClick={(e) => e.preventDefault()}
                          icon
                        >
                          <MoreVertical />
                        </Button>
                      </DropdownMenuTrigger>
                      <DropdownMenuContent className="w-40">
                        <DropdownMenuLabel>Edit</DropdownMenuLabel>
                        <DropdownMenuSeparator />
                        <DialogTrigger
                          onClick={() => {
                            setDeleteNode(file);
                          }}
                        >
                          <DropdownMenuItem>
                            <Trash className="mr-2 h-4 w-4" />
                            <span>Delete</span>
                          </DropdownMenuItem>
                        </DialogTrigger>
                        <DialogTrigger
                          onClick={() => {
                            setRenameNode(file);
                          }}
                        >
                          <DropdownMenuItem>
                            <TextCursorInput className="mr-2 h-4 w-4" />
                            <span>Rename</span>
                          </DropdownMenuItem>
                        </DialogTrigger>
                      </DropdownMenuContent>
                    </DropdownMenu>
                  </div>
                </div>
              );
            })}
            <DialogContent>
              {deleteNode && (
                <Delete
                  node={deleteNode}
                  close={() => {
                    setDialogOpen(false);
                  }}
                />
              )}
              {renameNode && (
                <Rename
                  node={renameNode}
                  close={() => {
                    setDialogOpen(false);
                  }}
                  unallowedNames={
                    data?.children?.results.map((x) => x.name) || []
                  }
                />
              )}
            </DialogContent>
          </Dialog>
        </div>
      </div>
    </div>
  );
};

export default ExplorerPage;
