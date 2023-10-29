import { createTRPCRouter, protectedProcedure } from "@merc/server/api/trpc";
import { z } from "zod";

export const workspaceRouter = createTRPCRouter({
  getAll: protectedProcedure.query(({ ctx }) => {
    return ctx.prisma.workspace.findMany({
      include: {
        _count: true,
      },
    });
  }),

  get: protectedProcedure
    .input(
      z.object({
        workspaceId: z.string(),
      }),
    )
    .query(({ ctx, input }) => {
      return ctx.prisma.workspace.findUnique({
        where: {
          id: input.workspaceId,
        },
      });
    }),

  create: protectedProcedure
    .input(
      z.object({ 
      }),
    )
    .mutation(({ ctx, input }) => {
      return ctx.prisma.workspace.create({
        data: { 
        },
      });
    }),

  update: protectedProcedure
    .input(
      z.object({
        workspaceId: z.string(), 
      }),
    )
    .mutation(({ ctx, input }) => {
      return ctx.prisma.workspace.update({
        where: {
          id: input.workspaceId, 
        },
      });
    }),

  remove: protectedProcedure
    .input(
      z.object({
        workspaceId: z.string(),
      }),
    )
    .mutation(async ({ ctx, input }) => {
      return ctx.prisma.workspace.delete({
        where: { id: input.workspaceId }
      });      
    })
});
