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
			  legal_name: z.string(),
			  name: z.string(),
			  size: z.number().nullable().optional(),
      }),
    )
    .mutation(({ ctx, input }) => {
      return ctx.prisma.workspace.create({
        data: { 
			    legal_name: input.legal_name,
			    name: input.name,
			    size: input.size,
        },
      });
    }),

  update: protectedProcedure
    .input(
      z.object({
        workspaceId: z.string(), 
			  legal_name: z.string().nullable(),
			  name: z.string().nullable(),
			  size: z.number().nullable(),
      }),
    )
    .mutation(({ ctx, input }) => {
      return ctx.prisma.workspace.update({
        where: {
          id: input.workspaceId, 
			    legal_name: input.legal_name,
			    name: input.name,
			    size: input.size,
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
